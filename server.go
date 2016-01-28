package main

import (
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	//Parse Data
	body, error := ioutil.ReadAll(r.Body) 
	fileName := r.Header["Filename"]
	deviceID := r.Header["Identifier"]
	dirPath  := fmt.Sprint("logs/", deviceID[0], "/")
	filePath := fmt.Sprint(dirPath, fileName[0]);

	if error != nil {
		http.Error(w, error.Error(), 500)
	}

	//Write File	
 	os.MkdirAll(dirPath, os.ModePerm)	
	out, error := os.Create(filePath);
	if error != nil {
		http.Error(w, error.Error(), 500)
	}	

	_, error = out.Write([]byte(body));
	if error != nil {
		http.Error(w, error.Error(), 500)
	}	
	
	fmt.Print("Recieved File\n")
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8000", nil)
}
