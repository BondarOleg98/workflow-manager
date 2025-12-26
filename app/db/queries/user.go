package queries

const SaveUserQuery = "INSERT INTO users (email, username, password, created_at) VALUES ($1, $2, $3, $4)"
const GetUserByEmailQuery = "SELECT * FROM users WHERE email = $1"
const GetUserByIdQuery = "SELECT id, email, username, password, created_at, last_login FROM users WHERE id = $1"
