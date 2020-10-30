INSERT INTO users
  (id, name, nickname)
VALUES
  (1, 'Tommy Stark', 'Man of iron'),
  (2, 'Peter Potter', 'The very best neighbor'),
  (3, 'Harry Parker', NULL),
  (4, 'Bayek of Siwa', 'Amun, the last Medjay');

INSERT INTO task_states
  (id, state)
VALUES
  (1, 'TODO'),
  (2, 'In progress'),
  (3, 'Blocked'),
  (4, 'Done');

INSERT INTO tasks
  (user_id, description, state_id)
VALUES
  (1, 'Build the new Bill-51 suit', 1),
  (2, 'Find a new hideout', 1),
  (3, 'Create a Nickname (wall of shame is getting me)', 3),
  (3, 'Bake the most delicious dessert (remember, low calories)', 2);