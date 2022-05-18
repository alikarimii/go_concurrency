## Confinement

> Confinement is the simple yet powerful idea of ensuring information is only ever available from one concurrent process
> When this is achieved, a concurrent program is implicitly safe and no synchronization is needed.
>
> There are two kinds of confinement possible:
>
> - ad hoc (base on convention with the group you work within or codebase)
> - lexical (compiler to enforce the confinement)
>
> When working with concurrent code, there are a few different options for safe operation:
>
> - Synchronization primitives for sharing memory (e.g., sync.Mutex)
> - Synchronization via communicating (e.g., channels)
> - Immutable data
> - Data protected by confinement
>   (Concurrency in Go - Katherine Cox-Buday)

---

## Preventing Goroutine Leaks

> The goroutine has a few paths to termination:
>
> - When it has completed its work.
> - When it cannot continue its work due to an unrecoverable error.
> - When it’s told to stop working.

---

## Or-Channel

> At times you may find yourself wanting to combine one or more done channels into a single done channel that closes if any of its component channels close.
> (Concurrency in Go - Katherine Cox-Buday)

---

## Pipeline

### [Monads](https://ericlippert.com/2013/02/21/monads-part-one)

A [monad](https://stackoverflow.com/questions/2704652/monad-in-plain-english-for-the-oop-programmer-with-no-fp-background) is an "amplifier" of types that obeys certain rules and which has certain operations provided.

An **amplifie** is something that increases the representational power of their **underlying** type

example of a commonly-used C# type that is **monadic** in nature

- Nullable<T> — represents a T that could be null
- Func<T> — represents a T that can be computed on demand
- Lazy<T> — represents a T that can be computed on demand once, then cached
- Task<T> — represents a T that is being computed asynchronously and will be available in the future, if it isn’t already
- IEnumerable<T> — represents an ordered, read-only sequence of zero or more Ts

A byte can be one of 256 values; that’s very useful but also very simple. By using generic types we can represent “an asynchronously-computed sequence of nullable bytes” very easily:

- Task<IEnumerable<Nullable<byte>>>. That adds a huge amount of power to the “byte” type without changing its fundamental “byte-ish” nature.

> **properties of a pipeline stage:**
>
> - A stage consumes and returns the same type.
> - A stage must be reified by the language so that it may be passed around. Func‐ tions in Go are reified and fit this purpose nicely.A stage consumes and returns the same type.
>   (Concurrency in Go - Katherine Cox-Buday)

### Pros and cons

If one of your stages is computationally expensive, this will certainly eclipse this performance overhead.

---

## Fan-In

A function can read from multiple inputs and proceed until all are closed by multiplexing the input channels onto a single channel that’s closed when all the inputs are closed. This is called fan-in.

> fanning in means multiplexing or joining together multiple streams of data into a single stream.
> (Concurrency in Go - Katherine Cox-Buday)

## Fan-Out

Multiple functions can read from the same channel until that channel is closed; this is called fan-out. This provides a way to distribute work amongst a group of workers to parallelize CPU use and I/O.

#### Notice

> A naive implementation of the fan-in, fan-out algorithm only works if the order in which results arrive is unimportant. We have done nothing to guarantee that the order in which items are read from the randIntStream is preserved as it makes its way through the sieve. Later, we’ll look at an example of a way to maintain order.
> (Concurrency in Go - Katherine Cox-Buday)

---

## tee-channel

> Sometimes you may want to split values coming in from a channel so that you can send them off into two separate areas of your codebase.
> Imagine a channel of user commands: you might want to take in a stream of user commands on a channel, send them to something that executes them, and also send them to something that logs the commands for later auditing

## Bridge channel

> In some circumstances, you may find yourself wanting to consume values from a sequence of channels: (channel of channels)
>
> - <-chan <-chan interface{}

## Queue

> Sometimes it’s useful to begin accepting work for your pipeline even though the pipeline is not yet ready for more. This process is called queuing.
> All this means is that once your stage has completed some work, it stores it in a tem‐ porary location in memory so that other stages can retrieve it later, and your stage doesn’t need to hold a reference to it

> In queuing theory, there is a law that with enough sampling predicts the throughput of your pipeline. It’s called Little’s Law
> It is commonly expressed as: L=λW, where:
>
> - L = the average number of units in the system.
> - λ = the average arrival rate of units.
> - W = the average time a unit spends in the system.
>
> This equation only applies to so-called **stable systems**. In a pipeline, a stable system is one in which the rate that work enters the pipeline, or **ingress**, is equal to the rate in which it exits the system, or **egress**

---

## Context

> the context package serves two primary purposes:
>
> - To provide an API for canceling branches of your call-graph. (Done)
> - To provide a data-bag for transporting request-scoped data through your call-graph.(Value)

> Go authors recommend you follow a few rules when storing and retrieving value from a Context:
>
> - define a custom key-type in your package
> - Use context values only for request-scoped data that **transits processes and API boundaries**, not for passing optional parameters to functions.

### transits processes and API boundaries:

> 1. The data should transit process or API boundaries.
>    If you generate the data in your process’ memory, it’s probably not a good candi‐ date to be request-scoped data unless you also pass it across an API boundary.
> 1. The data should be immutable.
>    If it’s not, then by definition what you’re storing did not come from the request.
> 1. The data should trend toward simple types.
>    If request-scoped data is meant to transit process and API boundaries, it’s much easier for the other side to pull this data out if it doesn’t also have to import a complex graph of packages.
> 1. The data should be data, not types with methods.
>    Operations are logic and belong on the things consuming this data.
> 1. The data should help decorate operations, not drive them.
>    If your algorithm behaves differently based on what is or isn’t included in its Context, you have likely crossed over into the territory of optional parameters.

| Data                  | 1   | 2   | 3   | 4   | 5   |
| --------------------- | --- | --- | --- | --- | --- |
| Request ID            | ✓   | ✓   | ✓   | ✓   | ✓   |
| UserID                | ✓   | ✓   | ✓   | ✓   |     |
| URL                   | ✓   | ✓   |     |     |     |
| API Server Connection |     |     |     |     |     |
| Authorization Token   | ✓   | ✓   | ✓   | ✓   |     |
| Request Token         | ✓   | ✓   | ✓   |     |     |

---

## Scatter/Gather

this pattern is base on this article in [medium](https://medium.com/star-gazers/go-scatter-gather-pattern-a439c70afe16)
