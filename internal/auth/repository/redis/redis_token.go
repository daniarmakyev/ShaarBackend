package redis

type RedisTokenRepository struct{}

func (r *RedisTokenRepository) SaveToken(userID int, token string) error {
	return nil
}
