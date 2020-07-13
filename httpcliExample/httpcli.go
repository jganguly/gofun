package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	myHttpClient()
}

func myHttpClient() {

	/* Ensure there is a TimeOut */
	var netClient = &http.Client{
		Timeout: time.Millisecond * 5000,
	}

	/* payload sent to the server */
	ses := Ses{ Key: "name", Val: "Jaideep"}
	requestBody, error := json.Marshal(ses)
	if error != nil {
		panic(error)
	}

	fmt.Println("key:val sent to server", ses)

	/* post call to server */
	response, error := netClient.Post("http://localhost:3001/setsesval","application/json", bytes.NewBuffer(requestBody))

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	/* response body from server */
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}



	/* post call to server */
	response, error = netClient.Post("http://localhost:3001/getsesval","application/json", bytes.NewBuffer(requestBody))

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	/* response body from server */
	body, error = ioutil.ReadAll(response.Body)
	if error != nil {
		panic(error)
	}

	json.Unmarshal(body,&ses)
	fmt.Println("key:val received from server", ses)

}
