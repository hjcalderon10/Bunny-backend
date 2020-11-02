package task

const (
	create_task         = "INSERT INTO tasks (title, description, user_id) values ($1, $2, $3) RETURNING id;"
	get_all_tasks       = "SELECT * FROM tasks;"
	get_all_task_states = "SELECT * FROM task_states;"
	get_task            = "SELECT * FROM tasks WHERE id = $1;"
	update_task         = "UPDATE tasks SET %s WHERE id = $1;"
	delete_task         = "DELETE FROM tasks WHERE id = $1;"
)
