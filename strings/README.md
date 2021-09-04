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


