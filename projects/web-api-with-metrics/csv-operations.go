package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

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

type WineSubset struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// Structures for converting csv to JSON
type State struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Ts     string `json:"timestamp"`
}

func readCSV(path *string) ([][]string, error) {
	initLog := log.New(os.Stdin, "ERROR\t", log.Ldate|log.Ltime)
	csvFile, err := os.Open(*path)
	if err != nil {
		initLog.Println("error open file", err)
		return nil, err
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	data, err := reader.ReadAll()
	if err != nil {
		initLog.Println("error read file", err)
		return nil, err
	}
	return data, nil
}

func getWineId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	data, _ := readCSV(&filePath)
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
		}
	}
	for _, item := range wines {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func getWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	start, _ := strconv.Atoi(r.URL.Query().Get("start"))
	count, _ := strconv.Atoi(r.URL.Query().Get("count"))
	data, _ := readCSV(&filePath)
	for i, line := range data {
		if i > 0 {
			wineslist = append(wineslist, WineSubset{
				Id:    line[0],
				Title: line[11],
			})
		}
	}
	if start == 0 && count == 0 {
		json.NewEncoder(w).Encode(wineslist)
	} else {
		json.NewEncoder(w).Encode(wineslist[start : count+1])
	}
}

func createWine(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var wine Wine
	_ = json.NewDecoder(r.Body).Decode(&wine)
	wine.Id = strconv.Itoa(idcounter)
	wines = append(wines, wine)
	idcounter++
	json.NewEncoder(w).Encode(wines)
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
