package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	var port, url string
	var exists bool

	if url, exists = os.LookupEnv("HTTP_ADDR"); !exists {
		url = "0.0.0.0"
	}

	if port, exists = os.LookupEnv("HTTP_PORT"); !exists {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", ApplicationDetails).Methods("Get")
	r.HandleFunc("/.well-known/live", Live).Methods("Get")
	r.HandleFunc("/.well-known/ready", Ready).Methods("Get")
	fmt.Printf("Service started at %s:%s\n", url, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", url, port), r))

}

func ApplicationDetails(w http.ResponseWriter, r *http.Request) {
	var backgroundColor, versionColor, envColor, environment string
	var exists bool
	if backgroundColor, exists = os.LookupEnv("BACKGROUND_COLOR"); !exists {
		backgroundColor = "white"
	}

	if versionColor, exists = os.LookupEnv("VERSION_COLOR"); !exists {
		versionColor = "black"
	}

	if envColor, exists = os.LookupEnv("ENV_COLOR"); !exists {
		envColor = "black"
	}

	if environment, exists = os.LookupEnv("ENV"); !exists {
		environment = "default"
	}

	version, err := ioutil.ReadFile("./VERSION")
	if err != nil {
		version = []byte("No version set")
	}

	htmlData := `<!DOCTYPE html>
	<html>
	<body style="background-color:` + backgroundColor + `;">
	<h2 style="color:` + versionColor + `;">Deployed Version: ` + string(version) + ` </h2>
	<h2 style="color:` + envColor + `;">Environment: ` + environment + `</h2>
	<h3 style="color:` + envColor + `;">List of all environment variables</h3>`

	envVars := os.Environ()
	sort.Strings(envVars)
	for _, ev := range envVars {
		htmlData += ev + "<br>"
	}

	htmlData += `</body>
	</html>`

	htmlDataInBytes := []byte(htmlData)
	fmt.Println("Retrieved application details")
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
	w.Write(htmlDataInBytes)
}

func Live(w http.ResponseWriter, r *http.Request) {
	var livenessResponse string
	var exists bool
	if livenessResponse, exists = os.LookupEnv("RESPONSE_CODE"); !exists {
		livenessResponse = "204"
	}

	response, err := strconv.Atoi(livenessResponse)
	if err != nil {
		fmt.Println("Couldn't connect to database")
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(response)
}

func Ready(w http.ResponseWriter, r *http.Request) {
	var readinessResponse string
	var exists bool
	if readinessResponse, exists = os.LookupEnv("RESPONSE_CODE"); !exists {
		readinessResponse = "204"
	}

	response, err := strconv.Atoi(readinessResponse)
	if err != nil {
		fmt.Println("Couldn't connect to database")
		w.WriteHeader(400)
		return
	}

	w.WriteHeader(response)
}
