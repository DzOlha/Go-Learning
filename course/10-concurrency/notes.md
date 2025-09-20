# Go Concurrency Basics

This document summarizes some key notes on using **goroutines**, **channels**, and the **defer** keyword in Go.

---

## Goroutines

A **goroutine** is a lightweight thread of execution.  
To start a goroutine, use the `go` keyword before the function call:

```go
go someFunctionCall()
```

> Goroutines do not return values directly. To communicate results, use **channels**.

---

## Channels

A **channel** is a typed conduit that allows goroutines to synchronize and share data.

Example of sending and receiving signals with a channel:

```go
func someFunc(channel chan bool) {
    // do some work...
    channel <- true // signal completion
}

doneChannel := make(chan bool)      // create a channel
go someFunc(doneChannel)            // start goroutine and pass channel
<- doneChannel                      // wait until goroutine finishes
```

---

## Select Statement

When working with multiple channels, you can use the `select` statement to wait on multiple channel operations simultaneously. It executes the first case that is ready:

```go
for index, value := range totalNumberOfItems {
    select {
    case v1 := <-ch1:
        if v1 != nil {
            fmt.Println("received from ch1")
        }
    case v2 := <-ch2:
        fmt.Println("received from ch2")
    }
}
```

---

## Parallel Execution with Channel Slices

You can create a slice of channels to manage multiple goroutines in parallel:

```go
dones := make([]chan bool, 4)

for index, value := range someMyArray {
    dones[index] = make(chan bool)
    go myFunc(dones[index])
}

// Wait for all goroutines to finish
for _, done := range dones {
    <-done
}
```

---

## The `defer` Keyword

The `defer` keyword schedules a function call to run **after the surrounding function finishes execution** (whether normally or due to a panic).

Example:

```go
defer file.Close() // will run at the end of the surrounding function
```

---

## Summary

- Use `go` to launch functions concurrently.
- Use **channels** to send signals or data between goroutines.
- Use `select` to listen on multiple channels at once.
- Use channel slices to manage many concurrent goroutines.
- Use `defer` to postpone execution until the end of the current scope.

This provides a foundation for writing safe and efficient concurrent programs in Go.
