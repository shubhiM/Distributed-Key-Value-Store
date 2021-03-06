Contributors:
    Shubhi Mittal
    Shraddha phadnis

In-memory Distributed key-value store
==============================================

In-memory distributed key-value store design consists of 3 parts namely Client, Server and Proxy/router

Client: Client is just a user who sends requests to set the key-values, fetch all the key-values in database, query a particular key.
curl commands are used to send respective requests. Below are few examples:

Set API(Request/Response):
---------------------------
curl -H 'Content-Type: application/json' -X PUT -d '[{"key": {"encoding": "binary","data": "1234"},"value": { "encoding": "binary","data": "1010101010"}},{"key": { "encoding": "string","data": "1234"},"value": {"encoding": "string","data": "abcdefasdasdag"}}, {"key": {"encoding": "binary","data": "8959"},"value": { "encoding": "binary","data": "8959"}},{"key": { "encoding": "string","data": "##1asda2324"},"value": {"encoding": "string","data": "abcdefasdasdag"}}]' http://localhost:8080/set

Fetch API(Request/Response):
----------------------------
curl -i -H "Accept: application/json" -X GET http://localhost:8080/fetch

Query API(Request/Response)
---------------------------
curl -i -d '[{"encoding":"string", "data":"1234"}, {"encoding":"string", "data":"1234"}]' -H "Accept: application/json" -X POST http://localhost:8080/query

Proxy:
======

Proxy/Router keeps track of available servers and data stored in those servers. Proxy is always run on port 8080.
Proxy is made aware of all the servers running by passing servers as command line arguments.
Functionalities:

1) Proxy includes support for 3 APIs(/set,/fetch,/query). Each API is handled using respective handlers.
2) Hash Function: In order to avoid single node having all the data, we have distributed data among all available nodes using hash code returned by hash function. Using availble nodes/servers and key data, hashfunction returns the hash value which is then trasnlated to appropriate node where request is sent.

How to run Proxy:
go run proxy.go <hostname>:<port> <hostname>:<port>

Server:
=======
The program design for Server includes three pieces.
1. server.py: This class implements Simple http server and inherits the basic http server class provided in python.
2. api.py: The server requests are routed to apis defined in api.py according to the route definition and corresponding handler gets called.
3. database.py: Database class defined in this module simulates the in memory key value database. It is a Singleton class which has only one
live object in memory per process.

For each server process running on machine we have one database instance shared across all http requests.

Instructions to run the program:
===============================
1) Run servers as per the requirement as follows:
python httpServer.py <hostname> <port>
for example to run 2 servers:
             python httpServer.py localhost 3000
			 python httpServer.py localhost 3001

2) Run the proxy:
go run proxy.go localhost:3000 localhost:3001

3) Execute the curl command mentioned in client explanation section to send request to server
