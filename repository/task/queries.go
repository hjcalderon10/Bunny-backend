package task

const (
	create_task   = "INSERT INTO tasks (title, description, user_id) values ($1, $2, $3);"
	get_all_tasks = "SELECT * FROM tasks;"
	get_task      = "SELECT * FROM tasks WHERE id = $1;"
	update_task   = "UPDATE tasks SET %s WHERE id = $1;"
	delete_task   = "DELETE FROM tasks WHERE id = $1;"
)
