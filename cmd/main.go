/*
|--------------------------------------------------------------------------
| Main
|--------------------------------------------------------------------------
|
| This is the entry point for listeners of the project.
| You can create and run goroutines for event listeners below before the HTTP listener.
|
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"gomora-dapp/interfaces/http/rest"
)

func init() {
	// load our environmental variables.
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	// rest port
	restPort, err := strconv.Atoi(os.Getenv("API_URL_REST_PORT"))
	if err != nil {
		log.Fatalf("[SERVER] Invalid port")
	}
	if len(fmt.Sprintf("%d", restPort)) == 0 {
		restPort = 8000 // default rest port is 8000 if not set
	}

	// serve rest server
	rest.ChiRouter().Serve(restPort)
}
