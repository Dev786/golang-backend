package sessionHelpers

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/satori/go.uuid"
)

const (
	REDIS_SERVER = "127.0.0.1:6379"
	TIME_TO_LIVE = 120
)

func GetSessionId() (string, error) {
	uuid, err := uuid.NewV4()
	return uuid.String(), err
}

func CreateRedisConnection() (redis.Conn, error) {
	connection, err := redis.DialURL("redis://127.0.0.1:6379")
	return connection, err
}

func AddSessionTokenToRedis(cache redis.Conn, tokenId string, username string) error {
	_, err := cache.Do("SETEX", tokenId, TIME_TO_LIVE, username)
	if err != nil {
		return err
	}
	return nil
}

func GetSessionToken(cache redis.Conn, sessionToken string) (string, error) {
	username, err := cache.Do("GET", sessionToken)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%q", username), nil
}

func CreateSession(username string) (string, error) {
	sessionId, _ := GetSessionId()
	redisConnection, _ := CreateRedisConnection()
	err := AddSessionTokenToRedis(redisConnection, sessionId, username)
	if err != nil {
		return "", err
	} else {
		return sessionId, nil
	}
}
