package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Greet string `json:"greeting"`
}

func (jsonResponse *Response) UnmarshalJSON(data []byte) error {
	var jsonRequest Request
	if err := json.Unmarshal(data, &jsonRequest); err != nil {
		return err
	}
	jsonResponse.Greet = "Hello, " + jsonRequest.Name + "!"
	return nil
}

func post_h(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := Request{}
	json.NewDecoder(r.Body).Decode(&u)
	uj, _ := json.Marshal(u)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	b := uj
	var m Response
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println(err)
		return
	}
	un, _ := json.Marshal(m)
	fmt.Fprintf(w, "%s", un)
}

func getuser(rw http.ResponseWriter, jsonRequest *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func main() {
	r := httprouter.New()
	r.GET("/hello/:name", getuser)
	r.POST("/hello", post_h)
	http.ListenAndServe("localhost:8080", r)
}