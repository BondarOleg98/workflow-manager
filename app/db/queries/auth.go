package queries

const SaveUserQuery = "INSERT INTO users (id, email, username, password, created_at) VALUES ($1, $2, $3, $4, $5)"
const GetUserByEmailQuery = "SELECT * FROM users WHERE email = $1"
const GetUserByIdQuery = "SELECT id, email, username, password, created_at, last_login FROM users WHERE id = $1"
const InsertRefreshTokenQuery = "INSERT INTO refresh_tokens (id, user_id, token, expired_at, created_at, revoked) VALUES ($1, $2, $3, $4, $5, $6)"
const GetRefreshTokenQuery = "SELECT id, user_id, token, expired_at, created_at, revoked FROM refresh_tokens WHERE token = $1"
const RevokedRefreshTokenQuery = "UPDATE refresh_tokens SET revoked = true WHERE token = $1"
