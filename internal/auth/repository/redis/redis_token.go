package redis

type RedisTokenRepository struct{}

func (r *RedisTokenRepository) SaveToken(userID int, token string) error {
	// Реализация работы с Redis
	return nil
}
