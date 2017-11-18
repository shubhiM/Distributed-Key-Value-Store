package main

import (
	"net/http"
	"flag"
	"fmt"
	"io/ioutil"
)

func main() {
	/*var (
	port = flag.String("p", "8080", "specify port to use defaulting to 8080");
	)*/
	flag.Parse();
	fmt.Printf("length of arguments = %d",len(flag.Args()));

	args := flag.Args()
	//fmt.Printf("args[1] = %d", args[1]);
	finish := make(chan bool)
	
	server1 := http.NewServeMux()
	server1.HandleFunc("/GET", getHandlerServer1)
	server1.HandleFunc("/post", postHandlerServer1)

	server2 := http.NewServeMux()
	server2.HandleFunc("/GET", getHandlerServer2)
	server2.HandleFunc("/post", postHandlerServer2)
	
	server3 := http.NewServeMux()
	server3.HandleFunc("/GET", getHandlerServer3)
	server3.HandleFunc("/post", postHandlerServer3)
	
	server4 := http.NewServeMux()
	server4.HandleFunc("/GET", getHandlerServer4)
	server4.HandleFunc("/post", postHandlerServer4)

	go func() {
		http.ListenAndServe(":"+args[0], server1)
	
	}()

	go func() {
		http.ListenAndServe(":"+args[1], server2)
	}()

	go func() {
		http.ListenAndServe(":"+args[2], server3)
	}()
	
	go func() {
		http.ListenAndServe(":"+args[3], server4)
	}()
	



	<-finish
}

var results[] string

/* get handlers definitions*/
func getHandlerServer1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server1 port:"))
}

func getHandlerServer2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server2 port:"))
}

func getHandlerServer3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server3:"))
}

func getHandlerServer4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server4:"))
}

/* post handlers */

func postHandlerServer1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server1: POST "))

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results,string(body))

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func postHandlerServer2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server2 port: POST "))

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results,string(body))

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func postHandlerServer3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server3 port: POST "))

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results,string(body))

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func postHandlerServer4(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listening on server4 port"))

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}
		results = append(results,string(body))

		fmt.Fprint(w, "POST done")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
