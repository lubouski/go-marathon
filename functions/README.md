# Functions
A Function declaration has four parts: the keyword `func`, the name of the function, the input parameters, and the return type. 

#### Variadic Unput Parameters and Slices
We've been using `fmt.Println` to print results to the screen and you've probably noticed that it allows any number of input parameters. How does it do that? Like many languages, Go supports `variadic parameters`. The variadic parameter must be the last (or only) parameter in the input parameters list. You indicate it with three dots (...) before the type. The variable that's created within the function is a `slice` of the specified type. You use it just like any other slice. 
```
  func addTo(base int, vals ...int) []int {
    out := make([]int, 0, len(vals))
    for _, v := range vals {
      out = append(out, base+v)
    }
    return out
  }
``` 
And now we'll call it a few different ways:
```
  func main() {
    fmt.Println(addTo(3)) \\ outputs []
    fmt.Println(addTo(3, 2)) \\ outputs [5]
    fmt.Println(addTo(3, 2, 4, 6, 8)) \\ outputs [5 7 9 11]
    a := []int{4, 3}
    fmt.Println(addTo(3, a...)) \\ outputs [7, 6]
    fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) \\ outputs [4 5 6 7 8]
  }
```

#### Multiple Return Values
Go allows for multiple return values.
```
  func divAndRemainder(numerator int, denominator int) (int, int, error) {
    if denominator == 0 {
      return 0, 0, errors.New("cannot divide by zero")
    }
    return numerator / denominator, numerator % denominator, nil
  }
```
Calling our function looks like this:
```
  func main() {
    result, reminder, err := divAndRemainder(5, 2)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    fmt.Println(result, remainder)
  }
```
#### Defined Type
Defined type supports the same operators which base type supports. 
i.e (+, -, \*, /, ==, >, <)



