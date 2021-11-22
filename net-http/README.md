# Networking, net/http 
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

