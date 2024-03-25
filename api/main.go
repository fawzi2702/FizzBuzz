package main

import (
	"github.com/fawzi2702/FizzBuzz/api/helpers/environment"
	"github.com/fawzi2702/FizzBuzz/api/server"
)

func main() {
	if envErr := environment.SetupEnv(); envErr != nil {
		panic(envErr)
	}

	srv, srvErr := server.SetupServer()
	if srvErr != nil {
		panic(srvErr)
	}

	if srvErr = server.StartServer(srv); srvErr != nil {
		server.ShutdownServer(srv)
		panic(srvErr)
	}
	defer server.ShutdownServer(srv)
}
