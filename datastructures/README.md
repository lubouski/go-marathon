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
