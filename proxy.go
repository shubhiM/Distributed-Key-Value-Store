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

type setResponse struct{
    keys_added string
    Keys_failed KeyVal `json:"Keys_failed"`
}

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


var servers [1000] string


func main() {
	flag.Parse();
	fmt.Printf("length of arguments = %d",len(flag.Args()));
    args := flag.Args()
	for i:=0; i< len(flag.Args());i++ {
         servers[i] = "http://"+ args[i]
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
    if r.Method != "PUT" {
        http.Error(w, "Only Post method is supported", 400)
    } else {
        request := make([]setObject,0)
        b, err1 := ioutil.ReadAll(r.Body)
        if err1 != nil {
            panic(err1)
            os.Exit(1)
        }
        defer r.Body.Close()
        json.Unmarshal(b, &request)
        for i := 0;i<len(request);i++ {
            serverNum := hashFunc(request[i].Key.Data) % (uint32)(len(flag.Args()))
            fmt.Println("Selected host via Hash in set", servers[serverNum])
            var requestArray [1] setObject
            requestArray[0] = request[i]
            fmt.Println("request/host", requestArray)
            output, err2 := json.Marshal(requestArray)

            fmt.Println("output", string(output))
            if err2 != nil {
                panic(err2)
                os.Exit(1)
            }
            newReq, err3 := http.NewRequest(
                "PUT",
                servers[serverNum] + "/set",
                bytes.NewBuffer(output))
            if err3 != nil {
                panic(err3)
                os.Exit(1)
            }
            newReq.Header.Set("Content-Type", "application/json")
            client := &http.Client{}
            resp, err4 := client.Do(newReq)
            if err4 != nil {
                panic(err4)
                os.Exit(1)
            }
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            // finalResponse := make([]setResponse,0)
            // json.Unmarshal(body, &finalResponse)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader((int)(resp.StatusCode))
            w.Write(body)
        }
    }
}

/* post handlers */
func fetchHandleFunc(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET"{
        for i := 0;i<len(flag.Args());i++ {
    	       response, err1 := http.Get(servers[i] + "/fetch")
               if err1 != nil {
                   panic(err1)
                   os.Exit(1)
                }
                    defer response.Body.Close()
                    contents, err2 := ioutil.ReadAll(response.Body)
                    if err2 != nil {
                        panic(err2)
                        os.Exit(1)
                    }
                    w.Header().Set("Content-Type","application/json")
                    w.WriteHeader((int)(response.StatusCode))
                    w.Write([]byte(contents))
        }
    }
    if r.Method == "POST" {
        request := make([]KeyVal, 0)
        b, err1 := ioutil.ReadAll(r.Body)
        if err1 != nil {
            panic(err1)
            os.Exit(1)
        }
        defer r.Body.Close()
        json.Unmarshal(b, &request)
        for i := 0;i<len(request);i++ {
            serverNum := hashFunc(request[i].Data) % (uint32)(len(flag.Args()))
            fmt.Println("Selected host via Hash in query", servers[serverNum])
            var requestArray [1] KeyVal
            requestArray[0] = request[i]
            fmt.Println("request/host", requestArray)
            output, err2 := json.Marshal(requestArray)
            if err2 != nil {
                panic(err2)
                os.Exit(1)
            }
            newReq, err3 := http.NewRequest(
                "POST",
                servers[serverNum]+"/fetch",
                bytes.NewBuffer(output))
            if err3 != nil {
                panic(err3)
                os.Exit(1)
            }
            newReq.Header.Set("Content-Type", "application/json")
            client := &http.Client{}
            resp, err4 := client.Do(newReq)
            if err4 != nil {
                panic(err4)
                os.Exit(1)
            }
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(resp.StatusCode)
            w.Write(body)
        }
    }
}

/* post handlers */
func queryHandleFunc(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Only Post method is supported", 400)
    } else {
        request := make([]KeyVal, 0)
        b, err1 := ioutil.ReadAll(r.Body)
        if err1 != nil {
            panic(err1)
            os.Exit(1)
        }
        defer r.Body.Close()
        json.Unmarshal(b, &request)
        for i := 0;i<len(request);i++ {
            serverNum := hashFunc(request[i].Data) % (uint32)(len(flag.Args()))
            fmt.Println("Selected host via Hash in query", servers[serverNum])
            var requestArray [1] KeyVal
            requestArray[0] = request[i]
            fmt.Println("request/host", requestArray)
            output, err2 := json.Marshal(requestArray)
            if err2 != nil {
                panic(err2)
                os.Exit(1)
            }
            newReq, err3 := http.NewRequest(
                "POST",
                servers[serverNum]+"/query",
                bytes.NewBuffer(output))
            if err3 != nil {
                panic(err3)
                os.Exit(1)
            }
            newReq.Header.Set("Content-Type", "application/json")
            client := &http.Client{}
            resp, err4 := client.Do(newReq)
            if err4 != nil {
                panic(err4)
                os.Exit(1)
            }
            defer resp.Body.Close()
            body, _ := ioutil.ReadAll(resp.Body)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(resp.StatusCode)
            w.Write(body)
        }
    }
}
