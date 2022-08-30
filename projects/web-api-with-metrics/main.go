package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	wines     []Wine
	wineslist []WineSubset
	idcounter int
	err       error
	filePath  string
)

// MidleWare wrapper logic part
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	buffer := flag.Int("buffer", 1000, "Channel buffer for non blocking number of reads once a minute collected metrics")
	flag.StringVar(&filePath, "filePath", "/tmp/go-marathon/projects/web-api-with-metrics/sample-data.csv", "Path to winemag csv file")
	metricsInterval := flag.Int("metricsInterval", 60, "service emit metrics interval")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// buffered channel for non blocking reads once a minute collected metrics
	c := make(chan int, *buffer)

	router := mux.NewRouter()
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/wine", getWines).Methods("GET")
	router.HandleFunc("/wine/{id}", getWineId).Methods("GET")
	router.HandleFunc("/wine", createWine).Methods("PUT")

	infoLog.Println("starting server on", *addr)
	wrapperdHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		router.ServeHTTP(w, r)
	})

	handler := wrapHandlerWithLogging(wrapperdHandler, c)

	go func() {
		log.Fatal(http.ListenAndServe(*addr, handler))
	}()
	for _ = range time.Tick(time.Duration(*metricsInterval) * time.Second) {
		drainchan(c)
	}
	//log.Fatal(http.ListenAndServe(":8000", r))
}

// MidleWare methods for WrapperHandler
func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler, c chan int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("--> %s %s", req.Method, req.URL.Path)

		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, req)

		statusCode := lrw.statusCode
		log.Printf("<-- %d %s", statusCode, http.StatusText(statusCode))
		if statusCode == 200 {
			c <- 1
		} else {
			c <- 0
		}
	})
}

// drainchan could help pereodically drain channel to sync router and logs emits go routines
func drainchan(commch chan int) {
	infoLog := log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime)
	var okSlice []int
	var errSlice []int
	for {
		select {
		case e := <-commch:
			if e == 1 {
				okSlice = append(okSlice, e)
			} else {
				errSlice = append(errSlice, e)
			}
		default:
			infoLog.Println("number of success req is:", len(okSlice))
			errLog.Println("number of error req is:", len(errSlice))
			infoLog.Println("number of wines is:", idcounter)
			if len(errSlice) == 0 && len(okSlice) == 0 {
				infoLog.Println("service was idle for last minute")
			} else {
				infoLog.Println("availability of the service for last minute:", (float64(len(okSlice))/float64((len(okSlice)+len(errSlice))))*float64(100), "%")
			}
			return
		}
	}
}
