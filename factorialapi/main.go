package main

import (
	"encoding/json"
	"factorialapi/database"
	"factorialapi/domain"
	"factorialapi/log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func setEnvironmentVariables() {
	env := os.Getenv("APP_ENV")
	var envFile string
	if env == "docker" {
		envFile = ".env.docker"
	} else {
		envFile = ".env.local"
	}
	godotenv.Load(envFile)
	log.Info("loaded environment variables")
}

func main() {
	setEnvironmentVariables()

	http.HandleFunc("/metric", handleMetrics)
	http.HandleFunc("/metric/", handleMetric)

	log.Info("starting server on :" + os.Getenv("WEBSERVERPORT"))
	http.ListenAndServe(":"+os.Getenv("WEBSERVERPORT"), nil)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("handleMetrics with " + r.Method + " verb")

	if r.Method == http.MethodGet {
		query := r.URL.Query()
		if query.Has("name") {
			getMetricsFilter(w, query.Get("name"))
		} else {
			getMetrics(w)
		}
	} else if r.Method == http.MethodPost {
		createMetric(w, r)
	} else {
		log.Error("Method " + r.Method + " not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleMetric(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Info("handleMetric with " + r.Method + " verb")
	log.Info("URL Path: " + r.URL.String())

	id := r.URL.Path[len("/metric/"):]

	if r.Method == http.MethodGet {
		getMetric(w, id)
	} else {
		log.Error("Method " + r.Method + " not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getMetrics(w http.ResponseWriter) {
	list, err := domain.Metric.List(domain.Metric{Collection: database.Collection[domain.Metric]{}})

	if err != nil {
		log.Error("Error fetching data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

func createMetric(w http.ResponseWriter, r *http.Request) {
	var metric domain.Metric
	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		log.Error("Error decoding POST body: " + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metric.Collection = database.Collection[domain.Metric]{}
	metric, err = domain.Metric.Add(metric)

	if err != nil {
		log.Error("Error adding data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(metric)
}

func getMetric(w http.ResponseWriter, id string) {
	metric, err := domain.Metric.Find(domain.Metric{Collection: database.Collection[domain.Metric]{}}, id)

	if err != nil {
		log.Error("Error fetching data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metric)
}

func getMetricsFilter(w http.ResponseWriter, name string) {
	metric := domain.Metric{Name: name, Collection: database.Collection[domain.Metric]{}}
	metrics, err := metric.FindList()

	if err != nil {
		log.Error("Error fetching data: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(metrics)
}
