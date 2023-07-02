# Go Routines and Channels in Golang 🚀

1. What are GoRoutines? What are Channels?
2. Unbuffered & Buffered Channels
3. The main difference b/w the above two kinds
4. Mujhe kya fayda? Where do I use Channels in development?

__________

### [1] What are GoRoutines? What are Channels?

_**GoRoutines:**_ A goroutine is a lightweight thread managed by the Go runtime. Think of it like a thread in Go but with less overhead than maintaining a thread in other languages.

Similar to how threads in Java communicate using BlockingQueue (or other similar thread-safe DS), GoRoutines communicates using Channels in Golang.

Think of Channels as a pipe that can be used to send and receive data. Any Goroutine can send data and any other goroutine can receive data using Channels.
```go
ch <- v  // Send v(any data like integer) to channel ch.
v := <-ch // Receive from ch, and assign value to v.
```
__________

### [2] Unbuffered & Buffered Channels

Unbuffered example:
```go
ch := make(chan int) // no size is mentioned, meaning only 1 send and 1 receive will be done at a time
```
Buffered Channels:
```go
ch := make(chan int, 100) // max 100 integers can be accommodated in the channel at any given time
```

__________

### [3] The main difference b/w the above two kinds

Unbuffered channels do not have a capacity and will BLOCK the sending goroutine until the receiving goroutine is ready to receive the data.

Buffered channels have a capacity and will not block the sending goroutine until the channel is full.

__________

### [4] Mujhe kya fayda? Where do I use Channels in development?

Channels can be used anywhere where you can do work using multiple Goroutines like —

- **Worker Pools**: Multiple identical workers process incoming tasks from a shared channel
- **Producer-Consumer**: Multiple producers publish to a channel and multiple consumers consume from the same channel
- **Fan out, fan in**: Multiple GoRoutines work and publish to a shared channel but another goroutine consumes the final results
- Graceful shutdown of your service so that no data in the channel is lost

__________

_Overall, it's pretty much the same, it's a smart way in which you can modify the communication among your GoRoutines effectively._ 👍
