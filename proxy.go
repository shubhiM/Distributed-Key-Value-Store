package main

import (
	"net/http"
	"flag"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"hash/fnv"
	"os"
	//"reflect"
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

type serverList struct {
	server int 

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
	fmt.Printf("%s", r.URL)
	switch r.Method {
		case "GET":
			http.Error(w, "GET not supported", 400)
		case "POST":
			http.Error(w, "POST not supported", 400)
    	case "PUT":
    		// allocation using make
    		//url := "http://localhost:3000/set"
    		request := make([]setObject,0)
    		decode := json.NewDecoder(r.Body)
    		err := decode.Decode(&request)
    		if err != nil {
    			http.Error(w,err.Error(),400)
    		}
    		
    		for i:=0; i<len(request); i++ {
    			//fmt.Printf("keyvalue = %d",request[i].key.data)
    			//server := hashFunc(request[i].key.data) % (uint32)(len(flag.Args()))
    			fmt.Println(r.Body)
    			/*json := ioutil.ReadAll(r.Body)
    			//var json = []byte (r.Body)
    			req, err := http.NewRequest("PUT", url, json)
    			req.Header.Set("Content-Type","application/json")
    			client := &http.Client{}
    			resp , err := client.Do(req)

    			fmt.Println("response :", resp.Status)
    			body, _ := ioutil.ReadAll(resp.Body)
    			fmt.Println("response body:", string(body))*/
    			
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
	response, err := http.Get("http://localhost:3000/fetch")
    if err != nil {
        fmt.Printf("%s", err)
        os.Exit(1)
    } else {
        defer response.Body.Close()
        contents, err := ioutil.ReadAll(response.Body)
        if err != nil {
            fmt.Printf("%s", err)
            os.Exit(1)
        }
        fmt.Printf("%s\n", string(contents))
        w.Header().Set("Content-Type","application/json")
        w.Write([]byte(contents))
	
	}
}

/* post handlers */

func queryHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("queryHandleFunc: query"))
}
