package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	/* HTTP Server */
	myHttpSever()
}

var session Ses

func myHttpSever() {

	/* Handle function */
	http.HandleFunc("/setsesval",
		func(response http.ResponseWriter, request *http.Request) {

			response.Header().Set("Content-Type", "text/html; charset=utf-8")
			response.Header().Set("Access-Control-Allow-Origin", "*")

			t1 := time.Now().Nanosecond()
			// Decode the payload from the client
			decoder := json.NewDecoder(request.Body)
			var ses Ses
			err := decoder.Decode(&ses)
			if err != nil {
				panic(err)
			}

			// Set session
			fmt.Println(true,ses.Key)
			fmt.Println(true,ses.Val)
			session = ses

			t2 := time.Now().Nanosecond()
			t := t2 - t1
			fmt.Println(true,t)
		})

	http.HandleFunc("/getsesval",
		func(response http.ResponseWriter, request *http.Request) {

			response.Header().Set("Content-Type", "text/html; charset=utf-8")
			response.Header().Set("Access-Control-Allow-Origin", "*")

			t1 := time.Now().Nanosecond()
			fmt.Println(true,"In getsesval")
			// Decode the payload from the client
			decoder := json.NewDecoder(request.Body)
			var ses Ses
			err := decoder.Decode(&ses)
			if err != nil {
				panic(err)
			}

			// get session
			fmt.Println(true,session.Key)
			fmt.Println(true,session.Val)
			response.Header().Set("Content-Type", "application/json")
			json.NewEncoder(response).Encode(session)

			t2 := time.Now().Nanosecond()
			t := t2 - t1
			fmt.Println(true,t)
		})



	/* Serving Static Pages - not required for mobile
	   usage: http://localhost:3001/dummy.txt
	 */
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	////http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./public")))

	/* Listen & Serve */
	http.ListenAndServe(":3001", nil)
}

