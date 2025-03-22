package main

func main() {
	server := &LoggingServer{}
	client := &LoggingClient{
		server:      server,
		BatchedLogs: []Log{},
		BatchSize:   10,
	}

	// write good logs relevant to Log Levels not hello world
	client.Log("Database connection established", INFO)
	client.Log("High memory usage detected: 85% of available memory", WARN)
	client.Log("Failed to connect to database: connection timeout", ERROR)
	client.Log("User authentication successful", INFO)
	client.Log("Cache cleared successfully", INFO)
	//more logs
	client.Log("User logged in", INFO)
	client.Log("User logged out", INFO)
	client.Log("User registered", INFO)
	client.Log("User logged in", INFO)
	client.Log("User logged out", INFO)
	client.Log("User registered", INFO)
	// warn
	client.Log("High memory usage detected: 85% of available memory", WARN)
	client.Log("Failed to connect to database: connection timeout", ERROR)
	client.Log("User authentication successful", INFO)
	client.Log("Cache cleared successfully", INFO)
	// error
	client.Log("Failed to connect to database: connection timeout", ERROR)
	client.Log("Payment processing failed: insufficient funds", ERROR)

}