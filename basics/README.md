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
	var i int   = 10
        var w int32 = 10
        if w == i {
                fmt.Println("equals")
        } else {
                fmt.Println("not equal")
        }

# command-line-arguments
./variables.go:16:7: invalid operation: w == i (mismatched types int32 and int)
```

Constants:
Constants can be `untyped`. This can be useful when working with numbers such as integer-type data. If the constant is untyped, it is explicitly converted, where typed constants are not.

Let's refer to code example `constants.go`:
If you declare a constant with a type, it will be that exact type. Here when we declare the constant leapYear, we define it as data type int32. Therefore it is a typed constant, which means it can only operate with int32 data types. The year constant we declare with no type, so it is considered untyped. Because of this, you can use it with any integer data type.

When hours was defined, it inferred that it was of type int because we did not explicitly give it a type, hours := 24. When we declared minutes, we explicitly declared it as an int32, minutes := int32(60). 


