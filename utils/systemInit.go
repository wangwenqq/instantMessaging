package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app", viper.Get("app"))
	fmt.Println("config mysql", viper.Get("mysql"))
	fmt.Println("config redis", viper.Get("redis"))

}

func InitRedis() {

	Redis = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	ctx := context.Background()
	pong, err := Redis.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(pong)
	}

}

func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})

}

const PublishKey = "websocket"

func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	err = Redis.Publish(ctx, channel, msg).Err()
	return err
}

func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Redis.PSubscribe(ctx, channel)
	fmt.Println("Subscribe...", sub)
	msg, err := sub.ReceiveMessage(ctx)
	return msg.Payload, err
}
