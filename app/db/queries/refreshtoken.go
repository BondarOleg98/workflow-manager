package queries

const InsertRefreshTokenQuery = "INSERT INTO refresh_tokens (user_id, token, expired_at, created_at, revoked) VALUES ($1, $2, $3, $4, $5)"
const GetRefreshTokenByValueQuery = "SELECT id, user_id, token, expired_at, created_at, revoked FROM refresh_tokens WHERE token = $1"
const GetRefreshTokenByUserIdQuery = "SELECT id, user_id, token, expired_at, created_at, revoked FROM refresh_tokens WHERE user_id = $1"
const RevokedRefreshTokenQuery = "UPDATE refresh_tokens SET revoked = true WHERE token = $1"
