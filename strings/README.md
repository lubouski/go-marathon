# Strings & Runes & Bytes
Under the covers, Go uses a sequence of bytes to represent a string. These bytes don't have to be in any particular character encoding.
But several Go library functions assume that a string is composed of a sequence of UTF-8-encoded code points.

Just like you can extract a single value from an array or a slice, you can extract a single value from a string by using index expression:
```
  var s string = "hello California"
  var b byte = s[6]
```

Go allows you to pass a string to the built-in `len` function to find the length of the string. 
```
  var s string = "Hello t"
  fmt.Pringln(len(s))
```

#### strings & strconv Packages
Package `strings` contains function TrimSpace which could delete every space symbols (new line, tabs, spaces)
```
  s := "\t formerly surrounded by space \n"
  fmt.Println(strings.TrimSpace(s))
```
We could convert string into float64 type
```
  input = strings.TrimSpace(input)
  x, err := strconv.ParseFloat(input, 64)
```
#### Labeling Your FOR Statements
By default, the `break` and `continue` keywords apply to the loop that directly contains them. What if you have nested for loops and you want to exit or skip over an iterator of an outer loop? Let's look at an example. We're going to modify our string and stop iterating through a string as soon as it hits a terre "l"
```
func main() {
	samples := []string{"hello", "apple_n!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}
```
