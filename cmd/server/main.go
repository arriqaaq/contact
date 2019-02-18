package main

import (
	"bitbucket.org/greedygame/contact/api"
	"fmt"
	lucio "github.com/arriqaaq/server"
	"github.com/arriqaaq/x"
	kitlog "github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
)

const (
	SERVER_PORT = 8080
	SERVER_HOST = "0.0.0.0"
)

var (
	fieldKeys = []string{"method", "error"}
)

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	serverURL := fmt.Sprintf("%s:%d", SERVER_HOST, SERVER_PORT)
	// postgresConnectionStr := "host=db" + " user=" + os.Getenv("POSTGRES_USER") +
	// 	" dbname=" + os.Getenv("POSTGRES_DB") + " sslmode=disable password=" + os.Getenv("POSTGRES_PASSWORD")
	postgresConnectionStr := "host=db" + " user=" + os.Getenv("POSTGRES_USER") +
		" sslmode=disable password=" + os.Getenv("POSTGRES_PASSWORD")
	fmt.Println("postgres connection string: ", postgresConnectionStr)
	storageDb, stErr := gorm.Open(
		"postgres", postgresConnectionStr,
	)
	x.FatalCheck(stErr, "postgres connection error")
	defer storageDb.Close()

	storageDb.LogMode(true)

	storageDb.DropTableIfExists(&api.Contact{}, &api.Book{})
	storageDb.AutoMigrate(&api.Book{}, &api.Contact{})
	storageDb.Model(&api.Contact{}).AddForeignKey("book_id", "books(id)", "CASCADE", "CASCADE")

	var logger kitlog.Logger
	logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stderr))
	logger = kitlog.With(logger, "ts", kitlog.DefaultTimestampUTC)

	var httpLogger kitlog.Logger
	httpLogger = kitlog.With(logger, "component", "http")

	var cs api.Service
	cs = api.NewService(storageDb, logger)
	cs = api.NewLoggingService(logger, cs)
	cs = api.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "service_contact_book",
			Name:      "total_request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "service_contact_book",
			Name:      "success_request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "service_contact_book",
			Name:      "request_latency",
			Help:      "Total duration of requests in seconds.",
		}, fieldKeys),
		cs,
	)

	mux := http.NewServeMux()
	mux.Handle("/v1/", api.MakeHandler(cs, httpLogger))
	mux.Handle("/metrics", promhttp.Handler())

	httpHandler := accessControl(mux)
	server := lucio.NewServer(httpHandler, SERVER_HOST, SERVER_PORT)

	logger.Log("transport", "http", "address", serverURL, "msg", "listening")
	err := server.Serve()
	logger.Log("terminated", err)
}
