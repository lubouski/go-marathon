## io package provides clean interface to working with data
Any data structure which has method `Read` is valid for `io.Reader` type. In the example `main.go` file, we have a buffer which could help us to read data in chunks and save some memory.

### Copy function
Copy function play key role in helping us working with buffer.
```
package main

import (
    "fmt"
)

func main() {
    s := "brown fox over lazy dog"
    buf := make([]byte, 8)
    n := copy(buf, s[0:])
    fmt.Println(string(buf))
    fmt.Println("n offset is", n)

    fmt.Println("Printing next 8 bytes")

    n = copy(buf, s[:n])
    fmt.Println(string(buf))
}
```
The key is `n` function output which repserents offset number for the whole slice.
```
$ go run main.go
brown fo
n offset is 8
Printing next 8 bytes
x over
```
