## Service API for winemag data
Generally service requirements could be devided into three parts: `convert csv to JSON`, `expose APIs and logic`, `emit metrics to Stderr`. To achieve these things we need to first read csv data and process it to the respective Golang struct, then to expose APIs we could use `gorillaMux` server (it was picked to the advanced functionality compare with standard library http server) and `HandleFunc` to precess bussiness logic to every API request. And last part to emit metrics to Stderr every minute, require to have a non blocking way of passing request information from router to special function which will be invoked every minute, it implemented with `buffered channel` and `produce/consumer pattern`, to run these logic every minute was used ticker and drainchunnel func to read all the data from the channel.

### How to use API service:
First we need to download dependencie:
```
$ go get github.com/gorilla/mux
```
Then we could run service with `go run` ro build our code to binary and run it. Additionally there are `flags` which could be provided to service.
```
$ go run main.go -filePath "your winemag csv path"
INFO	2022/08/22 12:23:31 starting server on :8080
2022/08/22 10:46:50 --> GET /status
2022/08/22 10:46:50 <-- 200 OK
2022/08/22 10:46:51 --> GET /status
2022/08/22 10:46:51 <-- 200 OK
2022/08/22 10:46:51 --> GET /status
2022/08/22 10:46:51 <-- 200 OK
2022/08/22 10:46:52 --> GET /status
2022/08/22 10:46:52 <-- 200 OK
2022/08/22 10:46:52 --> GET /status
2022/08/22 10:46:52 <-- 200 OK
INFO	2022/08/22 10:47:16 number of success req is: 5
ERROR	2022/08/22 10:47:16 number of error req is: 0
INFO	2022/08/22 10:47:16 number of wines is: 11
INFO	2022/08/22 10:47:16 availability of the service for last minute: 100 %
2022/08/22 10:47:40 --> PUT /wine
2022/08/22 10:47:40 <-- 200 OK
INFO	2022/08/22 10:47:46 number of success req is: 1
ERROR	2022/08/22 10:47:46 number of error req is: 0
INFO	2022/08/22 10:47:46 number of wines is: 12
INFO	2022/08/22 10:47:46 availability of the service for last minute: 100 %
```
To simulate above behaviour we could access `http://localhost:8080/status` in our respective browser and on separate terminal run `curl` `PUT` command to add wine and optionaly pipe to `jq`.
```
terminal2$ curl -s -X PUT -H "Content-Type: application/json" -d '{"contry":"spain","description":"nice wine","designation":"what the hell?","points":"96","price":"800$","province":"Lombardy","regionone":"1","regiontwo":"2","tastername":"Paul","tastertwitter":"@Paul55","title":"peach pounch","variety":"vary","winery":"yo olde mitre"}' localhost:8080/wine | jq
### omited output
...
 {
    "id": "11",
    "country": "",
    "description": "nice wine",
    "designation": "what the hell?",
    "points": "96",
    "price": "800$",
    "province": "Lombardy",
    "regionone": "1",
    "regiontwo": "2",
    "tastername": "Paul",
    "tastertwitter": "@Paul55",
    "title": "peach pounch",
    "variety": "vary",
    "winery": "yo olde mitre"
  }
]
```
### Areas of improvement
First of all, structure of code could be splited to different files component, refactor variables and code logic. Then it is possbile to improve and add http server advanced options for `timeout` for example. Probably there is a better way to implement emit of metrics operation every minute. State of variables between functions could be handled better.
For live production we could use additional counfiguration options for http server, may be rate limiting. CSV winemag could be downloaded from the web or from S3 bucket to then converted to JSON. It would be a great idea to add more logging for code, and improve comments.
