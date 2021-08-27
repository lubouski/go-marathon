# Golang Basics
To the roots! 

### Basic Types
Best source to gain some knowledge on basic types is https://tour.golang.org/basics/1 

So Go's basic types are:
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

As a strictly typed language, Go has a rich type system with a multitude of types to represent data of many forms.
```
	// code snippet of types mismatch
	var i int
        var w int32
        w = 10
        i = 10
        if w == i {
                fmt.Println("equals")
        } else {
                fmt.Println("not equal")
        }

# command-line-arguments
./variables.go:16:7: invalid operation: w == i (mismatched types int32 and int)
```
