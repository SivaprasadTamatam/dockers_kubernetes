package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func main() {
	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Set initial visits count
	client.Set(context.Background(), "visits", 0, 0)

	// Define HTTP route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Get current visits count from Redis
		visits, err := client.Get(context.Background(), "visits").Result()
		if err != nil {
			http.Error(w, "Error getting visits count", http.StatusInternalServerError)
			return
		}

		// Convert visits count to integer
		visitsInt, err := strconv.Atoi(visits)
		if err != nil {
			http.Error(w, "Error converting visits count", http.StatusInternalServerError)
			return
		}

		// Increment visits count
		client.Set(context.Background(), "visits", strconv.Itoa(visitsInt+1), 0)

		// Send response
		fmt.Fprintf(w, "Number of visits is %d", visitsInt)
	})

	// Start the web server
	fmt.Println("Listening on :8082...")
	http.ListenAndServe(":8082", nil)
}
