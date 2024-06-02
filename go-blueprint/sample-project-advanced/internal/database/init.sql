INSERT INTO tasks (id, name, description, is_completed) VALUES
('8fcd40c7-74d4-4481-95dd-b64f313c619e', 'Buy groceries', 'Milk, Eggs, Bread, Butter', FALSE),
('2d3a0401-c10d-4d62-915f-7a6747ddf2e1', 'Write report', 'Finish the quarterly report by Friday', FALSE),
('7a9b2fc8-191e-4dc8-8aae-45bb18f48ea5', 'Call plumber', 'Fix the leaking sink in the kitchen', TRUE),
('2dcfabb7-9758-4209-b116-5fea77c64064', 'Schedule meeting', 'Set up a meeting with the project team', FALSE),
('a9e53510-a4e7-453a-a1ea-3956bc8843bc', 'Pay bills', 'Pay electricity and water bills', FALSE) ON CONFLICT DO NOTHING;
