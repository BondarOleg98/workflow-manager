package queries

const InsertUserQuery = "INSERT INTO users (id, email, username, password, created_at) VALUES ($1, $2, $3, $4, $5)"
const SelectUserByEmail = "SELECT * FROM users WHERE email = $1"
const InsertRefreshToken = "INSERT INTO refresh_tokens (id, user_id, token, expired_at, created_at, revoked) VALUES ($1, $2, $3, $4, $5, $6)"
