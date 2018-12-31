package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Microsoft/go-eventgrid-handler/eventgrid"
)

// this is the structure for the data portion of event grid messages
// change this according to the message you're expecting
type person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// personHandler - handle event grid "person" messages
// w and r are standard http.Handler params
// msg is the event grid message that was parsed by the event grid handler
func personHandler(w http.ResponseWriter, r *http.Request, msg []eventgrid.Envelope) {

	if msg == nil || len(msg) == 0 {
		log.Println("ERROR", "Event Grid Message is empty")
		w.WriteHeader(500)
		return
	}

	var p person
	env := msg[0]

	// get the values from env.Data
	if err := json.Unmarshal(env.Data, &p); err != nil {
		w.WriteHeader(500)
		log.Println("ERROR:", err)
		return
	}

	log.Println("personHandler", env.ID)

	// TODO this is where you would process the "person message"

	// tell event grid message was processed successfully
	w.WriteHeader(200)
}
