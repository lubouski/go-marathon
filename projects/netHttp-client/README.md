## Standard Library HTTP Client 
Basically what code does it sends a `GET` request to `https://jsonplaceholder.typicode.com/todos/1`, created appropriate struct to decode json, and outputs a result to the console. 
```
There are functions in the net/http package to make GET,HEAD and POST calls. Avoid using these functions because they use the default client, which means they don't set a request timeout.
```
