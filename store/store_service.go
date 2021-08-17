package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

//Определим структуру для Redis клиента

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

const CacheDuration = 6 * time.Hour

//Инициализируем сервис хранения и возвращаем указатель на хранилище
func InitializedStore()*StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong,err :=redisClient.Ping(ctx).Result()
	if err != nil{
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.redisClient = redisClient
	return storeService
}

//Сохраним сопоставление между originalUrl и сгенерированным
func SaveUrlMapping(shortUrl string, originalUrl string, userId string)  {
	err:=storeService.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed to save URL | Error: %v - shortURL: %s - originalUrl: %s\n", err, shortUrl, originalUrl ))
	}
}

//Получим originalUrl для перенаправления с shortUrl
func GetInitialUrl(shortUrl string) string {
	result, err := storeService.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed GetInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}

