-- Seeds the 5 MVP zones. Parameters mirror script/Deploy.s.sol on the
-- contracts side; keep both in sync if any of these change.
INSERT INTO zones (id, name, lat, lon, threshold_mm, payout_per_day, max_days_per_week, active)
VALUES
    (1, 'Yogyakarta', -7.7956, 110.3695, 20, 25000000000000000000000, 3, TRUE),
    (2, 'Sleman',     -7.7326, 110.3752, 20, 25000000000000000000000, 3, TRUE),
    (3, 'Bantul',     -7.8827, 110.3288, 20, 25000000000000000000000, 3, TRUE),
    (4, 'Surakarta',  -7.5755, 110.8243, 20, 25000000000000000000000, 3, TRUE),
    (5, 'Semarang',   -6.9932, 110.4203, 20, 25000000000000000000000, 3, TRUE)
ON CONFLICT (id) DO UPDATE SET
    name              = EXCLUDED.name,
    lat               = EXCLUDED.lat,
    lon               = EXCLUDED.lon,
    threshold_mm      = EXCLUDED.threshold_mm,
    payout_per_day    = EXCLUDED.payout_per_day,
    max_days_per_week = EXCLUDED.max_days_per_week,
    active            = EXCLUDED.active;
