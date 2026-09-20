// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

import {Test} from "forge-std/Test.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {IDRP} from "../src/IDRP.sol";
import {PayungPool} from "../src/PayungPool.sol";

contract PayungPoolTest is Test {
    IDRP idrp;
    PayungPool pool;

    address admin = makeAddr("admin");
    address oracle = makeAddr("oracle");

    uint16 constant ZONE_ID = 1;
    uint256 constant PREMIUM_PER_WEEK = 6_500 ether;
    uint16 constant THRESHOLD_MM = 20;
    uint256 constant PAYOUT_PER_DAY = 25_000 ether;
    uint8 constant MAX_DAYS_PER_WEEK = 3;

    function setUp() public {
        vm.startPrank(admin);
        idrp = new IDRP(admin);
        pool = new PayungPool(address(idrp), admin);
        pool.grantRole(pool.ORACLE_ROLE(), oracle);
        pool.createZone(ZONE_ID, "Bantul", PREMIUM_PER_WEEK, THRESHOLD_MM, PAYOUT_PER_DAY, MAX_DAYS_PER_WEEK);
        vm.stopPrank();

        // Anchor block.timestamp to a stable midday WIB moment so today()+1 math
        // in tests is not sensitive to the wall clock the test runs at.
        vm.warp(1_800_000_000);
    }

    function _fundPool(uint256 amount) internal {
        vm.prank(admin);
        idrp.mint(admin, amount);
        vm.prank(admin);
        idrp.approve(address(pool), amount);
        vm.prank(admin);
        pool.fundPool(amount);
    }

    function _giveDriver(address driver, uint256 amount) internal {
        vm.prank(admin);
        idrp.mint(driver, amount);
        vm.prank(driver);
        idrp.approve(address(pool), type(uint256).max);
    }

    function _buyPolicy(address driver) internal returns (uint256 policyId) {
        _giveDriver(driver, PREMIUM_PER_WEEK);
        vm.prank(driver);
        policyId = pool.buyPolicy(ZONE_ID, 1);
    }

    // ------------------------------------------------------------------
    // 1. buy policy, premium moves to pool, event emitted
    // ------------------------------------------------------------------
    function test_BuyPolicy_MovesPremiumAndEmitsEvent() public {
        address driver = makeAddr("driver1");
        _giveDriver(driver, PREMIUM_PER_WEEK);

        uint32 expectedStart = pool.today() + 1;
        uint32 expectedEnd = expectedStart + 7;

        vm.expectEmit(true, true, true, true);
        emit PayungPool.PolicyBought(1, driver, ZONE_ID, expectedStart, expectedEnd, PREMIUM_PER_WEEK);

        vm.prank(driver);
        uint256 policyId = pool.buyPolicy(ZONE_ID, 1);

        assertEq(policyId, 1);
        assertEq(idrp.balanceOf(address(pool)), PREMIUM_PER_WEEK);
        assertEq(idrp.balanceOf(driver), 0);

        PayungPool.Policy memory p = pool.getPolicy(policyId);
        assertEq(p.holder, driver);
        assertEq(p.zoneId, ZONE_ID);
        assertEq(p.startDay, expectedStart);
        assertEq(p.endDay, expectedEnd);
        assertFalse(p.closed);
    }

    // ------------------------------------------------------------------
    // 2. buy second policy same zone reverts PolicyAlreadyActive
    // ------------------------------------------------------------------
    function test_BuyPolicy_SecondInSameZoneReverts() public {
        address driver = makeAddr("driver1");
        _buyPolicy(driver);

        _giveDriver(driver, PREMIUM_PER_WEEK);
        vm.prank(driver);
        vm.expectRevert(PayungPool.PolicyAlreadyActive.selector);
        pool.buyPolicy(ZONE_ID, 1);
    }

    // ------------------------------------------------------------------
    // 3. rain 15 mm (below 20): settle reverts NotRainDay
    // ------------------------------------------------------------------
    function test_Settle_BelowThresholdReverts() public {
        address driver = makeAddr("driver1");
        _buyPolicy(driver);
        _fundPool(PAYOUT_PER_DAY * 10);

        uint32 policyStartDay = pool.getPolicy(1).startDay;
        vm.warp(uint256(policyStartDay + 1) * 1 days - 7 hours);

        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, policyStartDay, 15);

        vm.expectRevert(PayungPool.NotRainDay.selector);
        pool.settle(ZONE_ID, policyStartDay);
    }

    // ------------------------------------------------------------------
    // 4. rain 25 mm: 3 holders paid, pool balance decreases by 3 x payout
    // ------------------------------------------------------------------
    function test_Settle_PaysAllEligibleHolders() public {
        address d1 = makeAddr("driver1");
        address d2 = makeAddr("driver2");
        address d3 = makeAddr("driver3");
        _buyPolicy(d1);
        _buyPolicy(d2);
        _buyPolicy(d3);
        _fundPool(PAYOUT_PER_DAY * 10);

        uint32 policyStartDay = pool.getPolicy(1).startDay;
        vm.warp(uint256(policyStartDay + 1) * 1 days - 7 hours);

        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, policyStartDay, 25);

        uint256 poolBalanceBefore = pool.poolBalance();
        pool.settle(ZONE_ID, policyStartDay);

        assertEq(idrp.balanceOf(d1), PAYOUT_PER_DAY);
        assertEq(idrp.balanceOf(d2), PAYOUT_PER_DAY);
        assertEq(idrp.balanceOf(d3), PAYOUT_PER_DAY);
        assertEq(pool.poolBalance(), poolBalanceBefore - 3 * PAYOUT_PER_DAY);
    }

    // ------------------------------------------------------------------
    // 5. same day settle twice: second call pays nobody
    // ------------------------------------------------------------------
    function test_Settle_TwiceSameDayPaysOnce() public {
        address driver = makeAddr("driver1");
        _buyPolicy(driver);
        _fundPool(PAYOUT_PER_DAY * 10);

        uint32 policyStartDay = pool.getPolicy(1).startDay;
        vm.warp(uint256(policyStartDay + 1) * 1 days - 7 hours);

        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, policyStartDay, 25);

        pool.settle(ZONE_ID, policyStartDay);
        assertEq(idrp.balanceOf(driver), PAYOUT_PER_DAY);

        pool.settle(ZONE_ID, policyStartDay);
        assertEq(idrp.balanceOf(driver), PAYOUT_PER_DAY);
    }

    // ------------------------------------------------------------------
    // 6. 4 rain days in one week: 4th day pays nobody (cap 3)
    // ------------------------------------------------------------------
    function test_Settle_CapsPayoutsPerWeek() public {
        address driver = makeAddr("driver1");
        _buyPolicy(driver);
        _fundPool(PAYOUT_PER_DAY * 10);

        uint32 startDay = pool.getPolicy(1).startDay;

        for (uint32 i = 0; i < 4; i++) {
            uint32 day = startDay + i;
            vm.warp(uint256(day + 1) * 1 days - 7 hours);
            vm.prank(oracle);
            pool.submitRainfall(ZONE_ID, day, 25);
            pool.settle(ZONE_ID, day);
        }

        assertEq(idrp.balanceOf(driver), PAYOUT_PER_DAY * MAX_DAYS_PER_WEEK);
    }

    // ------------------------------------------------------------------
    // 7. policy bought today is not paid for today (starts tomorrow)
    // ------------------------------------------------------------------
    function test_Settle_PolicyNotActiveOnPurchaseDay() public {
        address driver = makeAddr("driver1");
        uint32 purchaseDay = pool.today();
        _buyPolicy(driver);
        _fundPool(PAYOUT_PER_DAY * 10);

        // Report and settle rain for the purchase day itself (before policy starts).
        vm.warp(uint256(purchaseDay + 1) * 1 days - 7 hours);
        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, purchaseDay, 25);
        pool.settle(ZONE_ID, purchaseDay);

        assertEq(idrp.balanceOf(driver), 0);
    }

    // ------------------------------------------------------------------
    // 8. expired policy not paid and is pruned
    // ------------------------------------------------------------------
    function test_Settle_ExpiredPolicyNotPaidAndPruned() public {
        address driver = makeAddr("driver1");
        _buyPolicy(driver);
        _fundPool(PAYOUT_PER_DAY * 10);

        uint32 endDay = pool.getPolicy(1).endDay;

        // Warp well past the policy's end day, then report/settle rain for a day
        // at/after expiry.
        vm.warp(uint256(endDay + 2) * 1 days - 7 hours);
        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, endDay, 25);
        pool.settle(ZONE_ID, endDay);

        assertEq(idrp.balanceOf(driver), 0);
        assertEq(pool.getActivePolicies(ZONE_ID).length, 0);
        assertTrue(pool.getPolicy(1).closed);
        assertEq(pool.activePolicyOf(driver, ZONE_ID), 0);
    }

    // ------------------------------------------------------------------
    // 9. pool underfunded: PayoutSkipped emitted, no revert
    // ------------------------------------------------------------------
    function test_Settle_UnderfundedPoolSkipsPayout() public {
        address driver = makeAddr("driver1");
        uint256 policyId = _buyPolicy(driver);
        // Pool only holds the premium just paid in, far less than one payout.

        uint32 startDay = pool.getPolicy(policyId).startDay;
        vm.warp(uint256(startDay + 1) * 1 days - 7 hours);
        vm.prank(oracle);
        pool.submitRainfall(ZONE_ID, startDay, 25);

        vm.expectEmit(true, true, true, true);
        emit PayungPool.PayoutSkipped(policyId, startDay, "insufficient pool balance");
        pool.settle(ZONE_ID, startDay);

        assertEq(idrp.balanceOf(driver), 0);
    }

    // ------------------------------------------------------------------
    // 10. non-oracle calling submitRainfall reverts
    // ------------------------------------------------------------------
    function test_SubmitRainfall_NonOracleReverts() public {
        address stranger = makeAddr("stranger");
        uint32 yesterday = pool.today() - 1;
        bytes32 oracleRole = pool.ORACLE_ROLE();

        vm.prank(stranger);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, stranger, oracleRole)
        );
        pool.submitRainfall(ZONE_ID, yesterday, 25);
    }

    // ------------------------------------------------------------------
    // Fuzz: today() and week boundary math
    // ------------------------------------------------------------------
    function testFuzz_Today_MatchesReferenceFormula(uint32 unixTime) public {
        vm.assume(unixTime > 7 hours);
        vm.warp(unixTime);
        uint32 expected = uint32((uint256(unixTime) + 7 hours) / 1 days);
        assertEq(pool.today(), expected);
    }

    function testFuzz_WeekIndex_IsDayIndexDividedBySeven(uint32 dayIndex) public pure {
        uint32 week = dayIndex / 7;
        assertEq(week, dayIndex / 7);
        assertLe(week * 7, dayIndex);
        assertLt(dayIndex - week * 7, 7);
    }

    /// @notice Three timestamps around WIB midnight must land on the correct
    ///         day_index: just before midnight WIB, exactly at midnight WIB,
    ///         and just after. WIB is UTC+7, so midnight WIB is 17:00 UTC the
    ///         previous day.
    function test_Today_AroundWibMidnightBoundary() public {
        uint256 midnightWibUnix = 1_800_000_000 - (1_800_000_000 % 1 days) + 1 days - 7 hours;

        vm.warp(midnightWibUnix - 1);
        uint32 dayBefore = pool.today();

        vm.warp(midnightWibUnix);
        uint32 dayAt = pool.today();

        vm.warp(midnightWibUnix + 1);
        uint32 dayAfter = pool.today();

        assertEq(dayAt, dayBefore + 1);
        assertEq(dayAfter, dayAt);
    }
}
