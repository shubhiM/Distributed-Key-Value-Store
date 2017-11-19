package main

import (
	"net/http"
	"flag"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"hash/fnv"
	"os"
    "bytes"
)

type KeyVal struct {
	Encoding string `json:"encoding"`
	Data string `json:"data"`
}

type setObject struct {
	Key KeyVal `json:"key"`
	Value KeyVal `json:"value"`
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
	flag.Parse();
	fmt.Printf("length of arguments = %d",len(flag.Args()));
	for i:=0; i< len(flag.Args());i++ {

	}
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
	fmt.Printf("%s", r.URL)
	switch r.Method {
		case "GET":
			http.Error(w, "GET not supported", 400)
		case "POST":
			http.Error(w, "POST not supported", 400)
    	case "PUT":
    		request := make([]setObject,0)
            b, err := ioutil.ReadAll(r.Body)
	        defer r.Body.Close()
	        json.Unmarshal(b, &request)
	        output, err := json.Marshal(request)
            fmt.Println(err)
            newReq, err := http.NewRequest(
                    "PUT",
                    "http://localhost:3000/set",
                    bytes.NewBuffer(output))
            newReq.Header.Set("Content-Type", "application/json")
            client := &http.Client{}
            resp, err := client.Do(newReq)
            if err != nil {
                panic(err)
            }
            defer resp.Body.Close()
            fmt.Println("response Status:", resp.Status)
            fmt.Println("response Headers:", resp.Header)
            body, _ := ioutil.ReadAll(resp.Body)
            w.Header().Set("content-type", "application/json")
            w.Write(body)
    	case "DELETE":
    		http.Error(w, "DELETE not supported", 400)
    	default:
    		http.Error(w, "unknown request", 400)
		}
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
    if r.Method != "POST" {
        http.Error(w, "Only Post method is supported", 400)
    } else {
        w.Write([]byte("queryHandleFunc: query"))
        request := make([]KeyVal, 0)
        b, err := ioutil.ReadAll(r.Body)
        defer r.Body.Close()
        json.Unmarshal(b, &request)
        output, err := json.Marshal(request)
        fmt.Println(err)
        newReq, err := http.NewRequest(
                "POST",
                "http://localhost:3000/query",
                bytes.NewBuffer(output))
        newReq.Header.Set("Content-Type", "application/json")
        client := &http.Client{}
        resp, err := client.Do(newReq)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()
        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        w.Header().Set("content-type", "application/json")
        w.Write(body)
    }
}
