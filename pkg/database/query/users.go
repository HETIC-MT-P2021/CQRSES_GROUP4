package query

const (
	QUERY_FIND_USERS_BY_USERNAME string = "SELECT email, password, role FROM users WHERE Username = ?"
	QUERY_CREATE_ACCOUNT         string = "INSERT INTO app.users (Email, Username, Password, Role) VALUES (?, ?, ?, 1)"
)
