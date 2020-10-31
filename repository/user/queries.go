package user

const (
	create_user   = "INSERT INTO users (name, nickname, img_url) values ($1, $2, $3);"
	get_all_users = "SELECT * FROM users;"
	get_user      = "SELECT * FROM users WHERE id = $1;"
	update_user   = "UPDATE users SET %s WHERE id = $1;"
	delete_user   = "DELETE FROM users WHERE id = $1;"
)
