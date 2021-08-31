# Data structures Arrays & Slices & Maps
So lets start with arrays, the main reason why arrays exists in Go is to provide the backing store for slices,
which are one of the n=most useful features of Go.

#### Arrays
When using an array litteral to initialize an array, you can leave off the number and use **...** instead:
```
  var x = [...]int{10, 20, 30}
``` 

You can use **==** and **!=** to compare arrays:
```
  var x = [...]int{1, 2, 3}
  var y = [3]int{1, 2, 3}
  fmt.Println(x == y) // prints true
```

Go only has one-dimensional arrays, but you can simulate multidimensional arrays:
```
  var x [2][3]int
```

#### Slices
Slice is a type that isn't **comparable**. It is a compile-time error to use == to see if two slices are identical or != to see if they are different.
The only thing you can compare a slice with is **nil**:
```
  fmt.Println(x == nil) // prints true
```
The built-in **append** function is used to grow slices:
```
  var x = []int{1, 2, 3}
  x = append(x, 4, 5, 6)
```
One slice is appended onto another by using the **...** operator to expand source slice. 
```
  y := []int{20, 30, 40}
  x = append(x, y...)
```


