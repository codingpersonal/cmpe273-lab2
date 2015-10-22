package main

import (
		"net/http"
		"fmt"
		"io/ioutil"
		"encoding/json"
		"github.com/gorilla/mux"
)

// defines all the REST APIs handlers

// POST request handling
func PostHello(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)

	var req WelcomeMsgReq
	err = json.Unmarshal(body, &req)
    
    if err != nil {
    	fmt.Println("Error in parsing the body of the request");
    }  
    
    name := req.Name;
    var res WelcomeMsgRes
    res.Greeting = "Hello,"+name+"!"
	
	//change this res to the response which needs to be sent back
    json.NewEncoder(w).Encode(res)
}

//Get request handling
func GetHello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	 fmt.Fprintf(w, "Hello, %s!\n", name)
}

