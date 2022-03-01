## Run all test in this folder

```
go test 4.channels/*_test.go -v
```

## Channels

> Channels are one of the synchronization primitives in Go derived from Hoare’s CSP. While they can be used to synchronize access of the memory, they are best used to communicate information between goroutines
> (Concurrency in Go - Katherine Cox-Buday)

## Buffered Channels

> which are channels that are given a capacity when they’re instantiated. This means that even if no reads are performed on the channel, a goroutine can still perform n writes, where n is the capacity of the buffered channel
> (Concurrency in Go - Katherine Cox-Buday)

> `b := make(chan int, 4)`
> A buffered channel with no receivers and a capacity of four would be full after four writes, and block on the fifth write since it has nowhere else to place the fifth element.

> An unbuffered channel has a capacity of zero and so it’s already full before any writes

### Read from,Write to,and close nil channel

All will block

> Be sure to ensure the channels you’re working with are always initialized first

### Channel owner

> The goroutine that owns a channel should:
>
> - Instantiate the channel.
> - Perform writes, or pass ownership to another goroutine.
> - Close the channel.
> - Ecapsulate the previous three things in this list and expose them via a reader channel.
>   (Concurrency in Go - Katherine Cox-Buday)

### Channel Consumer

> As a consumer of a channel, I only have to worry about two things:
>
> - Knowing when a channel is closed.
> - Responsibly handling blocking for any reason.
>   (Concurrency in Go - Katherine Cox-Buday)
