package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Response struct { // Response  body to generate Success/Failure response
	Message string
	ID      string
	Status  int
}

var WorkQueue = make(chan WorkRequest, *queuesize) //Creating a work queue of size

func Collector(w http.ResponseWriter, r *http.Request) { //api to add each request(work) to Worker Queue

	//request body to json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var newentry WorkRequest
	json.Unmarshal(reqBody, &newentry)

	//validate
	validate := validator.New()
	err2 := validate.Struct(newentry)
	if err2 != nil {
		fmt.Println("error at validation", err2.Error())
		fmt.Fprintf(w, (string)(err2.Error()))
		return
	}

	//increment id ,assign id to the new entry
	id++
	newentry.ID = id

	//insert in worker queue
	WorkQueue <- newentry
	var response Response

	generateresponse(id, &response)

	json, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error at parsing response in redirect", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func generateresponse(ID int64, response *Response) {
	starttime := time.Now()
	for math.Abs((starttime.Sub(time.Now()).Seconds())) <= 60 {

		shared.mu.RLock()
		_, ok := shared.cache[ID]
		shared.mu.RUnlock()
		if ok {
			shared.mu.Lock()
			delete(shared.cache, ID)
			shared.mu.Unlock()
			response.Message = "Success"
			response.ID = strconv.FormatInt(int64(ID), 10)
			response.Status = 200
			return
		}
	}

	response.Message = "Failure,try again"
	response.ID = "__"
	response.Status = 500
	return
}
