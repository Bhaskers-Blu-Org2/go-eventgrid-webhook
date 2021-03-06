package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/Microsoft/go-eventgrid-handler/eventgrid"
	"github.com/Microsoft/go-eventgrid-handler/logb"
)

func TestMainFunc(t *testing.T) {
	go main()
	time.Sleep(500 * time.Millisecond)

	osChan <- os.Interrupt

	time.Sleep(500 * time.Millisecond)
}

func TestPersonHandler(t *testing.T) {

	e, err := genMessage(t)

	if err != nil {
		t.Error(err)
	}

	r, err := http.NewRequest("POST", "https://www.logb.com/person", bytes.NewBuffer(e)) // bytes.NewBufferString(s))
	if err != nil {
		t.Error("NewRequest: ", err)
	}

	w := httptest.NewRecorder()

	// chain the person handler
	h := http.Handler(logb.Handler(eventgrid.Handler(personHandler)))

	h.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Error("Return Code: ", w.Code)
	}

}

// genMessage - helper function to generate valid event grid json
func genMessage(t *testing.T) ([]byte, error) {
	env := eventgrid.Envelope{Subject: "person", EventType: "person", DataVersion: "1.0"}
	env.ID = "1001"
	env.EventTime = time.Now().UTC().Format("2006-01-02T15:04:05Z")

	p := person{FirstName: "John", LastName: "Doe"}

	var err error
	env.Data, err = json.Marshal(&p)

	if err != nil {
		return nil, err
	}

	var wrapper []eventgrid.Envelope
	wrapper = append(wrapper, env)

	return json.Marshal(wrapper)
}
