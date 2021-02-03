package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// structure of the object
type item struct {
	Key string `json:"key"`
	Value string `json:"value"`
	TimeStamp string
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	var itemList []item

	// GET request that response with a list of all objects inside the array
	r.Get("/list", func(rw http.ResponseWriter, r *http.Request) {
		// Parsse the array into json
		jData, _ := json.Marshal(itemList)

		// Response with the array in json form
		rw.Write(jData)
	})

	// GET request that adding new item object into the item object array
	r.Post("/add", func(rw http.ResponseWriter, r *http.Request) {
		// Read from the request body
		bodyContent, _ := ioutil.ReadAll(r.Body)
		
		// Parse the request body from json to an object and store it in the item object "newOne" 
		var newOne item
		json.Unmarshal([]byte(bodyContent), &newOne)

		// Add a timestamp
		newOne.TimeStamp = time.Now().Format(time.RFC850)

		// append item object array with object "newOne"
		itemList = append(itemList, newOne)

		// Form confirm
		rw.Write([]byte("Added!"))
	})

	// Server running on port 80
	http.ListenAndServe(":80", r)
}