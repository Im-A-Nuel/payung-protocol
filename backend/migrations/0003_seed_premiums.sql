-- Initial premium of Rp 5.000/week for every zone and month, per FR-01 in
-- docs/REQUIREMENTS.md. Without this row GET /zones would report Rp 0 while
-- PayungPool still charges the premium set at deploy time, so the driver
-- would be quoted a price the contract does not honour.
--
-- Phase 4 (POST /admin/pricing/recompute) overwrites every row here with a
-- figure derived from three years of rainfall, plus a real narrative.
INSERT INTO premiums (zone_id, month, expected_rain_days, premium_per_week, narrative_id, computed_at)
SELECT
    z.id,
    m.month,
    0,
    5000000000000000000000,
    'Premi awal Rp 5.000 per minggu. Angka ini belum dihitung dari riwayat hujan zona.',
    now()
FROM zones z
CROSS JOIN generate_series(1, 12) AS m (month)
ON CONFLICT (zone_id, month) DO NOTHING;
