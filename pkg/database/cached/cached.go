package cached

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go-jwt-template/internal/models"
	"go-jwt-template/pkg/config"
	"go.opentelemetry.io/otel/trace"
	"strconv"
	"time"
)

type SessionData struct {
	models.UserSession
	LoginTimeStamp time.Time `json:"login_time_stamp"`
}

type RedisSession struct {
	rdb               *redis.Client
	sessionExpiration time.Duration
}

type Session interface {
	Set(ctx context.Context, data SessionData) (string, error)
	Get(ctx context.Context, uuidKey string) (SessionData, error)
	GetUUID(ctx context.Context, userID string) (string, error)
	UpdateKey(ctx context.Context, uuidKeyOld string, userID int) (string, error)
	GetHSET(ctx context.Context, someKey string) (map[string]string, error)
	Delete(ctx context.Context, userID int, sessionID string) error
}

func InitRedis(tracer trace.Tracer) Session {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", viper.GetString(config.RedisHost), viper.GetInt(config.RedisPort)),
		Password: viper.GetString(config.RedisPassword),
		DB:       0,
	})

	if err := redisotel.InstrumentTracing(rdb); err != nil {
		panic(fmt.Sprintf("redis trace init: %v", err.Error()))
	}

	return RedisSession{rdb, time.Duration(viper.GetInt(config.SessionExpiration)) * time.Hour * 24}
}

func (r RedisSession) Set(ctx context.Context, data SessionData) (string, error) {
	sessionDataJSON, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	uuidBytes, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	key := uuidBytes.String()

	err = r.rdb.Set(ctx, key, sessionDataJSON, r.sessionExpiration).Err()
	if err != nil {
		return "", err
	}
	err = r.rdb.Set(ctx, strconv.Itoa(data.ID), key, r.sessionExpiration).Err()
	if err != nil {
		return "", err
	}

	return key, nil
}

func (r RedisSession) Get(ctx context.Context, uuidKey string) (SessionData, error) {
	rawJSONSessionData, err := r.rdb.Get(ctx, uuidKey).Result()
	if errors.Is(err, redis.Nil) {
		return SessionData{}, nil
	}
	if err != nil {
		err = fmt.Errorf("ошибка при чтении данных сессии из Redis: %v", err)
		return SessionData{}, err
	}

	var userSessionData SessionData

	err = json.Unmarshal([]byte(rawJSONSessionData), &userSessionData)
	if err != nil {
		return SessionData{}, err
	}

	return userSessionData, nil
}

func (r RedisSession) GetUUID(ctx context.Context, userID string) (string, error) {
	userUUID, err := r.rdb.Get(ctx, userID).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}
	if err != nil {
		err = fmt.Errorf("ошибка при чтении данных сессии из Redis: %v", err)
		return "", err
	}

	return userUUID, nil
}

func (r RedisSession) UpdateKey(ctx context.Context, uuidKeyOld string, userID int) (string, error) {
	newKeyBytes, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	newKey := newKeyBytes.String()

	_, err = r.rdb.Do(ctx, "RENAME", uuidKeyOld, newKey).Result()
	if err != nil {
		return "", err
	}
	_, err = r.rdb.Set(ctx, strconv.Itoa(userID), newKey, r.sessionExpiration).Result()
	if err != nil {
		return "", err
	}

	return newKey, nil
}

func (r RedisSession) GetHSET(ctx context.Context, someKey string) (map[string]string, error) {
	mapStringString, err := r.rdb.HGetAll(ctx, someKey).Result()
	if err != nil {
		return map[string]string{}, err
	}

	return mapStringString, nil
}

func (r RedisSession) Delete(ctx context.Context, userID int, sessionID string) error {
	result, err := r.rdb.Del(ctx, strconv.Itoa(userID)).Result()
	if err != nil {
		return err
	}

	result, err = r.rdb.Del(ctx, sessionID).Result()
	if err != nil {
		return err
	}

	// для примера
	if result == 0 {
		fmt.Println("Key does not exist")
	} else {
		fmt.Println("Key deleted")
	}

	return nil
}
