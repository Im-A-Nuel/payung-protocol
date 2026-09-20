// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {AccessControl} from "@openzeppelin/contracts/access/AccessControl.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

/// @title PayungPool
/// @notice Parametric rain insurance pool. Holds premiums and pushes payouts by rule;
///         the contract, not any operator, decides who gets paid.
contract PayungPool is AccessControl, ReentrancyGuard {
    using SafeERC20 for IERC20;

    bytes32 public constant ORACLE_ROLE = keccak256("ORACLE_ROLE");

    uint256 public constant MAX_ACTIVE_PER_ZONE = 50;

    struct Zone {
        string name;
        uint256 premiumPerWeek; // IDRP wei
        uint16 thresholdMm;
        uint256 payoutPerDay; // IDRP wei
        uint8 maxDaysPerWeek;
        bool active;
    }

    struct Policy {
        address holder;
        uint16 zoneId;
        uint32 startDay; // inclusive, day_index
        uint32 endDay; // exclusive
        bool closed;
    }

    IERC20 public immutable idrp;

    mapping(uint16 => Zone) public zones;
    mapping(uint256 => Policy) public policies;
    mapping(uint16 => uint256[]) internal activePolicyIds; // zoneId => list, max 50
    mapping(address => mapping(uint16 => uint256)) public activePolicyOf; // holder => zone => policyId (0 = none)
    mapping(uint16 => mapping(uint32 => uint16)) public rainfallMm; // zone => day => mm
    mapping(uint16 => mapping(uint32 => bool)) public rainReported;
    mapping(uint256 => mapping(uint32 => bool)) public paidForDay; // policyId => day => paid
    mapping(uint256 => mapping(uint32 => uint8)) public payoutsInWeek; // policyId => week => count

    uint256 public nextPolicyId = 1;

    event ZoneCreated(uint16 indexed zoneId, string name);
    event PremiumUpdated(uint16 indexed zoneId, uint256 premiumPerWeek);
    event PoolFunded(address indexed from, uint256 amount);
    event PolicyBought(
        uint256 indexed policyId,
        address indexed holder,
        uint16 indexed zoneId,
        uint32 startDay,
        uint32 endDay,
        uint256 premium
    );
    event RainfallReported(uint16 indexed zoneId, uint32 indexed dayIndex, uint16 mm, bool isRainDay);
    event PayoutSent(
        uint256 indexed policyId, address indexed holder, uint16 indexed zoneId, uint32 dayIndex, uint256 amount
    );
    event PayoutSkipped(uint256 indexed policyId, uint32 indexed dayIndex, string reason);
    event PolicyExpired(uint256 indexed policyId);

    error ZoneInactive();
    error PolicyAlreadyActive();
    error ZoneFull();
    error InvalidWeeks();
    error RainAlreadyReported();
    error DayNotFinished();
    error NotRainDay();
    error RainNotReported();

    constructor(address idrpToken, address admin) {
        idrp = IERC20(idrpToken);
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
    }

    // ---------------------------------------------------------------------
    // Admin
    // ---------------------------------------------------------------------

    function createZone(
        uint16 zoneId,
        string calldata name,
        uint256 premiumPerWeek,
        uint16 thresholdMm,
        uint256 payoutPerDay,
        uint8 maxDaysPerWeek
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        zones[zoneId] = Zone({
            name: name,
            premiumPerWeek: premiumPerWeek,
            thresholdMm: thresholdMm,
            payoutPerDay: payoutPerDay,
            maxDaysPerWeek: maxDaysPerWeek,
            active: true
        });
        emit ZoneCreated(zoneId, name);
    }

    function setPremium(uint16 zoneId, uint256 premiumPerWeek) external onlyRole(DEFAULT_ADMIN_ROLE) {
        zones[zoneId].premiumPerWeek = premiumPerWeek;
        emit PremiumUpdated(zoneId, premiumPerWeek);
    }

    function setZoneActive(uint16 zoneId, bool active) external onlyRole(DEFAULT_ADMIN_ROLE) {
        zones[zoneId].active = active;
    }

    // ---------------------------------------------------------------------
    // Pool funding
    // ---------------------------------------------------------------------

    function fundPool(uint256 amount) external {
        idrp.safeTransferFrom(msg.sender, address(this), amount);
        emit PoolFunded(msg.sender, amount);
    }

    // ---------------------------------------------------------------------
    // Policy lifecycle
    // ---------------------------------------------------------------------

    function buyPolicy(uint16 zoneId, uint8 weeksCount) external nonReentrant returns (uint256 policyId) {
        Zone storage zone = zones[zoneId];
        if (!zone.active) revert ZoneInactive();
        if (weeksCount < 1 || weeksCount > 4) revert InvalidWeeks();
        if (activePolicyOf[msg.sender][zoneId] != 0) revert PolicyAlreadyActive();
        if (activePolicyIds[zoneId].length >= MAX_ACTIVE_PER_ZONE) revert ZoneFull();

        uint256 premium = zone.premiumPerWeek * weeksCount;
        uint32 startDay = today() + 1;
        uint32 endDay = startDay + 7 * uint32(weeksCount);

        policyId = nextPolicyId++;
        policies[policyId] =
            Policy({holder: msg.sender, zoneId: zoneId, startDay: startDay, endDay: endDay, closed: false});
        activePolicyIds[zoneId].push(policyId);
        activePolicyOf[msg.sender][zoneId] = policyId;

        idrp.safeTransferFrom(msg.sender, address(this), premium);

        emit PolicyBought(policyId, msg.sender, zoneId, startDay, endDay, premium);
    }

    function pruneExpired(uint16 zoneId) external {
        uint32 today_ = today();
        uint256[] storage ids = activePolicyIds[zoneId];
        uint256 i = 0;
        while (i < ids.length) {
            uint256 pid = ids[i];
            Policy storage p = policies[pid];
            if (p.endDay <= today_) {
                _expirePolicyAt(ids, i, pid, p, zoneId);
            } else {
                i++;
            }
        }
    }

    // ---------------------------------------------------------------------
    // Oracle
    // ---------------------------------------------------------------------

    function submitRainfall(uint16 zoneId, uint32 dayIndex, uint16 mm) external onlyRole(ORACLE_ROLE) {
        if (dayIndex >= today()) revert DayNotFinished();
        if (rainReported[zoneId][dayIndex]) revert RainAlreadyReported();

        rainReported[zoneId][dayIndex] = true;
        rainfallMm[zoneId][dayIndex] = mm;

        bool isRainDay = mm >= zones[zoneId].thresholdMm;
        emit RainfallReported(zoneId, dayIndex, mm, isRainDay);
    }

    // ---------------------------------------------------------------------
    // Settlement
    // ---------------------------------------------------------------------

    function settle(uint16 zoneId, uint32 dayIndex) external nonReentrant {
        if (!rainReported[zoneId][dayIndex]) revert RainNotReported();

        Zone storage zone = zones[zoneId];
        if (rainfallMm[zoneId][dayIndex] < zone.thresholdMm) revert NotRainDay();

        uint32 today_ = today();
        uint32 week = dayIndex / 7;
        uint256[] storage ids = activePolicyIds[zoneId];

        uint256 i = 0;
        while (i < ids.length) {
            uint256 pid = ids[i];
            Policy storage p = policies[pid];

            bool activeOnDay = p.startDay <= dayIndex && dayIndex < p.endDay;
            if (activeOnDay && !paidForDay[pid][dayIndex] && payoutsInWeek[pid][week] < zone.maxDaysPerWeek) {
                if (idrp.balanceOf(address(this)) < zone.payoutPerDay) {
                    emit PayoutSkipped(pid, dayIndex, "insufficient pool balance");
                } else {
                    paidForDay[pid][dayIndex] = true;
                    payoutsInWeek[pid][week] += 1;
                    idrp.safeTransfer(p.holder, zone.payoutPerDay);
                    emit PayoutSent(pid, p.holder, zoneId, dayIndex, zone.payoutPerDay);
                }
            }

            if (p.endDay <= today_) {
                _expirePolicyAt(ids, i, pid, p, zoneId);
            } else {
                i++;
            }
        }
    }

    function _expirePolicyAt(uint256[] storage ids, uint256 index, uint256 pid, Policy storage p, uint16 zoneId)
        internal
    {
        uint256 last = ids.length - 1;
        if (index != last) {
            ids[index] = ids[last];
        }
        ids.pop();
        activePolicyOf[p.holder][zoneId] = 0;
        p.closed = true;
        emit PolicyExpired(pid);
    }

    // ---------------------------------------------------------------------
    // Views
    // ---------------------------------------------------------------------

    function today() public view returns (uint32) {
        return uint32((block.timestamp + 7 hours) / 1 days);
    }

    function getActivePolicies(uint16 zoneId) external view returns (uint256[] memory) {
        return activePolicyIds[zoneId];
    }

    function getPolicy(uint256 id) external view returns (Policy memory) {
        return policies[id];
    }

    function poolBalance() external view returns (uint256) {
        return idrp.balanceOf(address(this));
    }
}
