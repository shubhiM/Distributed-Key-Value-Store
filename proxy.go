package main

import (
	"net/http"
	"flag"
	"fmt"
	//"io/ioutil"
	"encoding/json"
	"hash/fnv"
	//"container/list"
)

type KeyVal struct {
	encoding string `json:"encoding"`
	data string `json:"data"`
}

type setObject struct {
	key KeyVal `json:"key"`
	value KeyVal `json:"value"`
}

func hashFunc(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32() 
}

var ServerList [1000][1000] setObject
func main() {
	/*var (
	port = flag.String("p", "8080", "specify port to use defaulting to 8080");
	)*/
	//length := len(flag.Args())
	flag.Parse();

	fmt.Printf("length of arguments = %d",len(flag.Args()));

	for i:=0; i< len(flag.Args());i++ {

	}
	// args := flag.Args()
	//fmt.Printf("args[1] = %d", args[1]);
	finish := make(chan bool)
	
	proxyserver := http.NewServeMux()
	proxyserver.HandleFunc("/set", setHandleFunc)
	proxyserver.HandleFunc("/fetch", fetchHandleFunc)
	proxyserver.HandleFunc("/query", queryHandleFunc)

	go func() {
		http.ListenAndServe(":8080", proxyserver)
	
	}()
	
	<-finish
}

var results[] string

/* get handlers definitions*/
func setHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Inside setHandleFunc:"))
	switch r.Method {
		case "GET":
			http.Error(w, "GET not supported", 400)
		case "POST":
			http.Error(w, "POST not supported", 400)
    	case "PUT":
    		// allocation using make
    		request := make([]setObject,0)
    		decode := json.NewDecoder(r.Body)
    		err := decode.Decode(&request)
    		if err != nil {
    			http.Error(w,err.Error(),400)
    		}
    	
    		for i:=0; i<len(request); i++ {
    			//fmt.Printf("keyvalue = %d",request[i].key.data)
    			server := hashFunc(request[i].key.data) % (uint32)(len(flag.Args()))
    			ServerList[server][i] = request[i]
    		}
    		//TBD
    	case "DELETE":
    		http.Error(w, "DELETE not supported", 400)
    	default: 
    		http.Error(w, "unknown request", 400)		
		}
		w.Write([]byte("Inside setHandleFunc ***:"))
}

/* post handlers */

func fetchHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fetchHandleFunc: fetch"))
}

/* post handlers */

func queryHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("queryHandleFunc: query"))
}
