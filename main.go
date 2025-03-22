package main

import (
	"fmt"
	"github.com/aman1117/logging-service/client"
	"github.com/aman1117/logging-service/server"
	"github.com/aman1117/logging-service/types"
)

func main() {
	server := server.NewLoggingServer()
	client := client.NewLoggingClient(server, 3)

	// System startup logs
	client.Log("Logging service initialized successfully", types.INFO)
	client.Log("Database connection established", types.INFO)
	
	// Warning logs for potential issues
	client.Log("High memory usage detected: 85% of available memory", types.WARN)
	client.Log("API response time increased by 200ms", types.WARN)
	
	// Error logs for critical issues
	client.Log("Failed to connect to database: connection timeout", types.ERROR)
	client.Log("Payment processing failed: insufficient funds", types.ERROR)
	
	// More info logs
	client.Log("User authentication successful", types.INFO)
	client.Log("Cache cleared successfully", types.INFO)

	// Force flush any remaining logs
	client.Flush()

	// Filter and display logs by level
	fmt.Println("\n=== INFO Logs ===")
	infoLogs := server.FilterBasedOnLevel(types.INFO)
	for _, log := range infoLogs {
		fmt.Printf("[%s] %s\n", log.Timestamp.Format("2006-01-02 15:04:05"), log.Message)
	}

	fmt.Println("\n=== WARN Logs ===")
	warnLogs := server.FilterBasedOnLevel(types.WARN)
	for _, log := range warnLogs {
		fmt.Printf("[%s] %s\n", log.Timestamp.Format("2006-01-02 15:04:05"), log.Message)
	}

	fmt.Println("\n=== ERROR Logs ===")
	errorLogs := server.FilterBasedOnLevel(types.ERROR)
	for _, log := range errorLogs {
		fmt.Printf("[%s] %s\n", log.Timestamp.Format("2006-01-02 15:04:05"), log.Message)
	}
}