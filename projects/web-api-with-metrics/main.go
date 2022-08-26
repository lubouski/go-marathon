package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var wines []Wine
var wineslist []Wines
var idcounter int
var err error

// MidleWare wrapper logic part
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// Structures for converting csv to JSON
type State struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Ts     string `json:"timestamp"`
}

type Wine struct {
	Id            string `json:"id"`
	Contry        string `json:"country"`
	Description   string `json:"description"`
	Designation   string `json:"designation"`
	Points        string `json:"points"`
	Price         string `json:"price"`
	Province      string `json:"province"`
	Regionone     string `json:"regionone"`
	Regiontwo     string `json:"regiontwo"`
	Tastername    string `json:"tastername"`
	Tastertwitter string `json:"tastertwitter"`
	Title         string `json:"title"`
	Variety       string `json:"variety"`
	Winery        string `json:"winery"`
}

type Wines struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	buffer := flag.Int("buffer", 1000, "Channel buffer for non blocking number of reads once a minute collected metrics")
	filePath := flag.String("filePath", "/tmp/vinemag/sample-data.csv", "Path to winemag csv file")
	metricsInterval := flag.Int("metricsInterval", 60, "service emit metrics interval")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// buffered channel for non blocking reads once a minute collected metrics
	c := make(chan int, *buffer)
	// call for initVines function
	wines, wineslist, idcounter, err = initVines(filePath)
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

func getStatus(w http.ResponseWriter, r *http.Request) {
	var st State
	ts := time.Now().Format(time.RFC3339)
	if err != nil {
		st = State{Status: "error", Msg: err.Error(), Ts: ts}
		json.NewEncoder(w).Encode(st)
	} else {
		st = State{Status: "ok", Msg: "", Ts: ts}
		json.NewEncoder(w).Encode(st)
	}
}

func getWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	if start == 0 && count == 0 {
		json.NewEncoder(w).Encode(wineslist)
	} else {
		json.NewEncoder(w).Encode(wineslist[start : count+1])
	}
}

func getWineId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range wines {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createWine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wine Wine
	var shortwine Wines
	_ = json.NewDecoder(r.Body).Decode(&wine)
	wine.Id = strconv.Itoa(idcounter)
	wines = append(wines, wine)
	shortwine.Id = strconv.Itoa(idcounter)
	shortwine.Title = wine.Title
	wineslist = append(wineslist, shortwine)
	idcounter++
	json.NewEncoder(w).Encode(wines)
}

func initVines(path *string) ([]Wine, []Wines, int, error) {
	initLog := log.New(os.Stdin, "ERROR\t", log.Ldate|log.Ltime)
	csvFile, err := os.Open(*path)
	if err != nil {
		initLog.Println("error open file", err)
		return nil, nil, 0, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		initLog.Println("error read file", err)
		return nil, nil, 0, err
	}
	for i, line := range data {
		if i > 0 {
			wines = append(wines, Wine{
				Id:            line[0],
				Contry:        line[1],
				Description:   line[2],
				Designation:   line[3],
				Points:        line[4],
				Price:         line[5],
				Province:      line[6],
				Regionone:     line[7],
				Regiontwo:     line[8],
				Tastername:    line[9],
				Tastertwitter: line[10],
				Title:         line[11],
				Variety:       line[12],
				Winery:        line[13],
			})
			wineslist = append(wineslist, Wines{
				Id:    line[0],
				Title: line[11],
			})
			idcounter++
		}
	}
	return wines, wineslist, idcounter, nil
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
