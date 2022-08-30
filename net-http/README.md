## Networking, net/http 
Package net/http help us to build simple Web applications.

```
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello, web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
```  
`writer` - is used to renew response to browser. `request` - is used to receive a request from browser.

Example server.go could be used as reference to request fields:
```
GET / HTTP/1.1
Header["Sec-Fetch-Site"] = ["none"]
Header["Sec-Fetch-User"] = ["?1"]
Header["Accept-Language"] = ["en-GB,en-US;q=0.9,en;q=0.8"]
Header["Connection"] = ["keep-alive"]
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"]
Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36"]
Header["Accept-Encoding"] = ["gzip, deflate, br"]
Header["Sec-Ch-Ua"] = ["\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"96\", \"Google Chrome\";v=\"96\""]
Header["Sec-Ch-Ua-Mobile"] = ["?0"]
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["Sec-Ch-Ua-Platform"] = ["\"macOS\""]
Header["Sec-Fetch-Mode"] = ["navigate"]
Header["Sec-Fetch-Dest"] = ["document"]
Host = "localhost:8000"
RemoteAddr = "127.0.0.1:54968"
``` 
### Pass arguments to HTTP handlers
An http.HandlerFunc is defined as:
```
type HandlerFunc func(ResponseWriter, *Request)
```
As you can see an http.HandlerFunc doesn’t take any arguments except an http.ResponseWriter and a pointer to an http.Request. In this article, we’re going to cover ways to pass data, configuration, and dependencies to these functions anyways.
* Global Variable
* Wrapped Handler
* Sruct with Handler

#### Global Variable
A straightforward way to pass data to an HTTP handler is through global variables. We just declare a global variable and use it in the handler definition:
```
package main

import "net/http"

const name = "UserService"

func main() {
    http.HandleFunc("/", testHandler)
    http.ListenAndServe(":8080", nil)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hi, from Service: " + name))
}
```
By using global variables is seldom the best approach, as they can easily be changed and used throughout the whole package or program (if you export them).

#### Wrapped Handler
A commonly used pattern is to define a wrapper function that takes some arguments and returns an http.HandlerFunc. The real handler function can then be declared as a closure and can access the arguments passed to the wrapping function:
```
package main

import "net/http"

func main() {
    http.HandleFunc("/", testHandler("UserService"))
    http.ListenAndServe(":8080", nil)
}

func testHandler(name string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hi, from Service: " + name))
    }
}
```
The wrapper function approach is good if there are just a few arguments you want to pass to the handler, and you don’t want to hold an internal state between requests.

#### Struct with Handler
The object-oriented approach to pass data is by internal object state. To do this in Go you can define a struct that has fields for data or dependencies and implements http.HandlerFunc methods, which access those fields:
```
import "net/http"

func main() {
    handlers := wrapperStruct{name: "UserService"}
    http.HandleFunc("/", handlers.testHandler)
    http.ListenAndServe(":8080", nil)
}

type wrapperStruct struct {
    name string
}

func (ws wrapperStruct) testHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hi, from Service: " + ws.name))
}
```
Make sure access to the fields is thread-safe as each incoming HTTP Request will execute the handler in a new goroutine.
