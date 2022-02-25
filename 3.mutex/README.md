## Mutex and RWMutex

---

### Mutex

share memory between multiple concurrent processes

### RWMutex

There may be memory that needs to be shared between multiple concurrent processes, but perhaps not all of these processes will read and write to this memory

## Cond

> “event” is any arbitrary signal between two or more
> goroutines that carries no information other than the
> fact that it has occurred. Very often you’ll want to wait
> for one of these signals before continuing execution on a goroutine.
>
> **we would some kind of way for a goroutine to efficiently sleep until it was signaled to wake and check its condition**
>
> (Concurrency in Go - Katherine Cox-Buday)

## Once

> Once is a type that utilizes some sync primitives internally to ensure that only one call to Do ever calls the function passed in **even on different goroutines**
> (Concurrency in Go - Katherine Cox-Buday)

## Pool

> As we’ve seen, the object pool design pattern is best used either when you have con‐ current processes that require objects, but dispose of them very rapidly after instantiation, or when construction of these objects could negatively impact memory
> (Concurrency in Go - Katherine Cox-Buday)
