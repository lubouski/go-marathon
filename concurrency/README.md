## Concurrency in golang
Any function can be launched as a goroutine. Goroutines are lightweight processes managed by the Go runtime. 

Goroutines communicate using channels. Like slices and maps, channels are a built-in type created using the make function:
```
 ch := nake(chan int)
```

#### Reading, Writing and Buffering
Use the <- operator to interact with a channel. 
```
  a := <-ch // reads a value from ch and assigns it to a
  ch <- b // write the value in b to ch
```

By default channels are unbuffered. Every write ti an open, unbuffered channel causes the writting goroutine to pause until another goroutine reads from the same channel. 

Go also has buffered channels. These channels buffer a limited number of writes without blocking. A buffered channel is created by specifying of the buffer when creating the channel:

```
  ch := make(chan int, 10)
``` 

#### for-range and Channels
You can also read from a channel using a for-range loop:
```
  for v := range ch {
    fmt.Println(v)
  }
```

#### Closing Channels
When you are done writting to a channel you close it using the built-in close function:
```
  close(ch)
```
Once a channel is closed, any attempts to write to the channel or close the channel again will `panic`.
In go we use the comma ok idiom to detect whether a channel has been closed or not:
```
  v, ok := <-ch
```
