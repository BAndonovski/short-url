package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BAndonovski/short-url/api"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func router() *mux.Router {
	r := mux.NewRouter()
	api.Router(r)
	return r
}

func checkSetting(setting string) {
	if len(os.Getenv(setting)) == 0 {
		panic(fmt.Sprintf("Mandatory setting not configured: %s", setting))
	}
}

func checkCrucialSettings() {
	checkSetting("FINN_PORT")
	checkSetting("FINN_PREFIX")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("Error loading .env file")
	}
	checkCrucialSettings()

	log.Fatal(http.ListenAndServe(os.Getenv("FINN_PORT"), router()))
}
