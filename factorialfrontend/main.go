package main

import (
	"factorialfrontend/api"
	"factorialfrontend/domain"
	"factorialfrontend/log"
	"html/template"
	"net/http"
	"os"
	"time"

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

	http.HandleFunc("/", index)
	http.HandleFunc("/metric", metric)
	http.HandleFunc("/metric/filter", metricFiltered)

	log.Info("starting server on :" + os.Getenv("WEBSERVERPORT"))
	http.ListenAndServe(":"+os.Getenv("WEBSERVERPORT"), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Info("handling index with " + r.Method + " verb")
	if r.Method == http.MethodGet {
		handleTemplate("resources/index.html", nil, w)
	} else {
		log.Error("Method " + r.Method + " not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func metricFiltered(w http.ResponseWriter, r *http.Request) {
	log.Info("handling metricFiltered with " + r.Method + " verb")
	if r.Method == http.MethodGet {
		listMetricsFiltered(w, r)
	} else {
		log.Error("Method " + r.Method + " not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func metric(w http.ResponseWriter, r *http.Request) {
	log.Info("handling metric with " + r.Method + " verb")
	if r.Method == http.MethodPost {
		add(w, r)
		listMetricsFiltered(w, r)
	} else if r.Method == http.MethodGet {
		listMetrics(w)
	} else {
		log.Error("Method " + r.Method + " not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func listMetricsFiltered(w http.ResponseWriter, r *http.Request) {
	filterTime := r.URL.Query().Get("filterTime")
	index, err := domain.NewIndexFiltered(api.Api[domain.Metric]{}, filterTime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error fetching data: " + err.Error())
		return
	}
	handleTemplate("resources/metrics.html", index, w)
}

func listMetrics(w http.ResponseWriter) {
	index, err := domain.NewIndex(api.Api[domain.Metric]{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error fetching data: " + err.Error())
		return
	}
	handleTemplate("resources/metrics.html", index, w)
}

func add(w http.ResponseWriter, r *http.Request) {
	timestamp := time.Now().UTC()
	name := r.FormValue("name")
	value := r.FormValue("value")

	metric := domain.Metric{
		Timestamp: timestamp, Name: name, Value: value, API: api.Api[domain.Metric]{},
	}
	_, err := metric.Add()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error adding data: " + err.Error())
		return
	}
}

func handleTemplate(filepath string, data any, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error("Error parsing template file: " + err.Error())
		return
	}
	tmpl.Execute(w, data)
}
