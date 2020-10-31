CREATE TABLE users
(
  id SERIAL,
  name VARCHAR NOT NULL,
  nickname VARCHAR,
  img VARCHAR,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  inactivated_at TIMESTAMP,
  CONSTRAINT users_pk PRIMARY KEY (id)
);

CREATE INDEX idx_users_name ON users(name);

CREATE TABLE task_states
(
  id SERIAL,
  state VARCHAR NOT NULL UNIQUE,
  CONSTRAINT task_states_pk PRIMARY KEY(id)
);

CREATE INDEX idx_tas_states_name ON users(state);


CREATE TABLE tasks
(
  id SERIAL,
  user_id INT NOT NULL,
  state_id INT NOT NULL DEFAULT 1,
  title VARCHAR NOT NULL DEFAULT 'New task',
  description TEXT DEFAULT 'What should I do now?',
  deadline TIMESTAMP,
  is_active BOOLEAN DEFAULT true,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  inactivated_at TIMESTAMP,
  CONSTRAINT tasks_pk PRIMARY KEY(id),
  FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
  FOREIGN KEY(state_id) REFERENCES task_states(id) ON DELETE CASCADE
);

CREATE INDEX idx_task_user_id ON tasks(user_id);