package main

import (
	"github.com/fawzi2702/FizzBuzz/api/helpers/environment"
	redisdb "github.com/fawzi2702/FizzBuzz/api/helpers/redis-db"
	"github.com/fawzi2702/FizzBuzz/api/server"
)

func main() {
	// Setup environment
	if envErr := environment.SetupEnv(); envErr != nil {
		panic(envErr)
	}

	// Setup RedisDB
	if rdbErr := redisdb.SetupRedisClient(); rdbErr != nil {
		panic(rdbErr)
	}

	// Setup server
	srv, srvErr := server.SetupServer()
	if srvErr != nil {
		panic(srvErr)
	}

	// Start server
	if srvErr = server.StartServer(srv); srvErr != nil {
		server.ShutdownServer(srv)
		panic(srvErr)
	}
	// Shutdown server
	defer server.ShutdownServer(srv)
}
