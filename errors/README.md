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

Simple errors handling could be done with `fmt.Errorf` or with `errors`. At the example below we could see that function return two value one of which is `fmt.Errorf("weight can't be negative, value privided %.2f", weight)`, without formating we could use errors.New("value can't be negative") as return value.
```
package main

import (
	"fmt"
	"log"
)

func main() {
	paint, err := paintNeeded(4.3, -2.8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("area for paint: %.2f\n", paint)
}

func paintNeeded(weight float64, height float64) (float64, error) {
	if weight < 0 {
		return 0, fmt.Errorf("weight can't be negative, value privided %.2f", weight)
	}
        if height < 0 {
                return 0, fmt.Errorf("height can't be negative, value privided %.2f", height)
        }
	return  weight * height, nil
}
``` 
