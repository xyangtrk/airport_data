package redisclient

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	"github.com/xyangtrk/airport_data/pkg/models"
)

// # local redis
// REDIS_HOST=localhost
// REDIS_PASSWORD=
// REDIS_DB=0
// REDIS_PORT=6379

var ctx = context.Background()

func InitializeRedisClient() *redis.Client {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get Redis connection details from environment variables
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	// redisDB := os.Getenv("REDIS_DB")

	// Convert redisDB to integer
	// redisDBInt, err := strconv.Atoi(redisDB)

	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		// DB:       redisDBInt,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	})
	return rdb
}

func InsertIntoRedis(outputMap map[string]models.OutputAirport, rdb *redis.Client) error {
	for code, airport := range outputMap {
		key := fmt.Sprintf("airports:%s", code)
		err := rdb.HSet(ctx, key, map[string]interface{}{
			"utcOffset":          airport.UTCOffset,
			"lat":                airport.Latitude,
			"lon":                airport.Longitude,
			"timeZoneRegionName": airport.TimeZoneRegionName,
		}).Err()

		if err != nil {
			return fmt.Errorf("failed to insert %s into Redis: %v", code, err)
		}
		log.Printf("Inserted %s into Redis", code)
	}

	log.Println("All airports inserted into Redis successfully")
	return nil
}
