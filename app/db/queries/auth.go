package queries

const InsertUserQuery = "INSERT INTO users (id, email, username, password, created_at) VALUES ($1, $2, $3, $4, $5)"
const SelectUserByEmail = "SELECT FROM users WHERE email = $1"
