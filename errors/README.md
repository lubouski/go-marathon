## Errors in GO
Go presents two standard ways of creating errors in standard library.
* errors.New 
* fmt.Errorf

```
type error interface {
  Error() string
}
```

Custom error example:
```
package main

import (
  "errors"
  "fmt"
  "os"
)

type RequestError struct {
  StatusCode int
  
  Err error
}

func (r *RequestError) Error() string {
  return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func doRequest() error {
  return &RequestError{
    StatusCode: 503,
    Err: errors.New("unavailable"),
  }
}

func main() {
  err := doRequest()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Println("success!")
}
```

`fmt` package will call `Error()` method automatically.
