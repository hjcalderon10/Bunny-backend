INSERT INTO users
  (name, nickname)
VALUES
  ('Tommy Stark', 'Man of iron'),
  ('Peter Potter', 'The very best neighbor'),
  ('Harry Parker', NULL),
  ('Bayek of Siwa', 'Amun, the last Medjay');

INSERT INTO task_states
  (state)
VALUES
  ('TODO'),
  ('In progress'),
  ('Blocked'),
  ('Done');

INSERT INTO tasks
  (user_id, description, state_id)
VALUES
  (1, 'Build the new Bill-51 suit', 1),
  (2, 'Find a new hideout', 1),
  (3, 'Create a Nickname (wall of shame is getting me)', 3),
  (3, 'Bake the most delicious dessert (remember, low calories)', 2);