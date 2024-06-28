INSERT INTO weather (name, temperature, humidity, wind_speed, condition)
VALUES
    ('Tashkent', 30, 50.5, 10, 'Sunny'),
    ('Samarkand', 28, 55.2, 8, 'Cloudy'),
    ('Bukhara', 33, 45.3, 12, 'Clear'),
    ('Khiva', 35, 40.1, 15, 'Hot'),
    ('Andijan', 27, 60.0, 5, 'Rainy')
;

INSERT INTO transport (number, stations, current_station, is_full)
VALUES
    (1, ARRAY['Station A', 'Station B', 'Station C'], 'Station A', FALSE),
    (2, ARRAY['Station D', 'Station E', 'Station F'], 'Station E', TRUE),
    (3, ARRAY['Station G', 'Station H', 'Station I'], 'Station G', FALSE),
    (4, ARRAY['Station J', 'Station K', 'Station L'], 'Station K', TRUE),
    (5, ARRAY['Station M', 'Station N', 'Station O'], 'Station M', FALSE)
;
