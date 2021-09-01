# Data structures Arrays & Slices & Maps
So lets start with arrays, the main reason why arrays exists in Go is to provide the backing store for slices,
which are one of the n=most useful features of Go.

### Arrays
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

### Slices
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
Built-in **make** function could create a slice with capacity and length which we want.
```
  x := make([]int, 0, 10)
  x = append(x, 5, 6, 7, 8)
```
In this case, we have a non-nil slice with a lenght of 0, but a capacity of 10. Slice the lenght is 0, we can't directly index into it, but we can append values to it.

#### copy
If you need to create a slice that's independent of the original, use the built-in copy function.
```
  x := []int{1, 2, 3, 4}
  y := make([]int, 4)
  num := copy(y, x)
  fmt.Println(y, num)
```
You could also copy from the middle of the slice:
```
  x := []int{1, 2, 3, 4}
  y := make([]int, 2)
  copy(y, x[2:])
```
The **copy** function allows you to copy between two slices that cover overlapping sections of an underlying slice:
```
  x := []int{1, 2, 3, 4}
  num = copy(x[:3], x[1:])
  fmt.Println(x, num)
```
In this case we are copyinh the last three values in x on top of the first three values of x, This prints out [2 3 4 4] 3.

### Maps
The map type is written as map[keyType]valueType. 
Let's take a look at a few ways to declare maps.
Declare a nilMap, the zero value for a map is **nil**.
```
  var nilMap map[string]int
```
We can use an empty map literal this is not the same as a nil map. It has a lenght of 0 but you can read and write to a map assigned an empty map literal.
```
  totalWins := map[string]int{}
```
And non empty map.
```
  teams := map[string][]string {
    "Orcas": []string{"Fred", "Ralph", "Bijou"},
    "Lions": []string{"Sarah", "Peter", "Billie"},
    "Kittens": []string{"Waldo", "Raul", "Ze"},
  }
```

Maps are like slices in several ways:
* Maps automatically grow as you add key-value pairs to them.
* If you know how many key-value pairs you plan to insert into a map, you can use **make** to create a map with a specific initial size.
* Passing a map to the **len** function tells you the number of key-value pairs in a map.
* The zero value for a map is **nil**
* Maps are not comparable. You can check if they are equal to nil.
* The key for a map can be any comparable type. This means you cannot use a slice or a map as the key for a map.

