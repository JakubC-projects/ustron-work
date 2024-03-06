INSERT INTO rounds(id, start_date, end_date, freeze_start_date) VALUES 
    (1, '2024-01-01', '2024-03-06 21:00:00+01:00', '2024-03-06 18:00:00+01:00'),
    (2, '2024-03-06 21:00:00+01:00', '2024-05-06- 21:00:00+01:00', '2024-05-06 18:00:00+01:00');

INSERT INTO persons (person_id, display_name, birth_date, team) VALUES 
    (54512, 'Test User Green', '2000-11-10', 'Green'),
    (54513, 'Test User Red', '2000-11-10', 'Red'),
    (54514, 'Test User Blue', '2000-11-10', 'Blue');


INSERT INTO on_track (team, status, round_id) VALUES 
    ('Blue', 0, 1),
    ('Green', 0, 1),
    ('Red', 0, 1),
    ('Orange', 0, 1),
    ('Female', 0, 1),
    ('Male', 0, 1);
