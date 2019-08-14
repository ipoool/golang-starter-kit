package drivers

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Redis - Redis driver connection
type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
	client   *redis.Client
}

// Connect - Connecting to the Redis Server
func (r *Redis) Connect() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.Host, r.Port),
		Password: r.Password,
		DB:       r.DB,
	})
}

// Ping - Check redis connection
func (r *Redis) Ping() bool {
	_, err := r.client.Ping().Result()
	if err != nil {
		return false
	}
	return true
}
