# Summary

Golang package for processing items using a worker queue model. The main point is to only spawn as many workers as there are CPUs available. The tasks to be completed will be added to a queue and as each worker completes a task, it checks the queue for remaining items.

This was *heavily* influenced by:

[Golang: Master/Worker in Go](http://devcry.heiho.net/2012/07/golang-masterworker-in-go.html)

My main goal was to make it as reusable as possible.

# Usage

You simply need to require the worker package: (Don't forget to run "go get" to fetch the external package)

```go
import "github.com/clarke/golang-worker"
```

And then write a function that accepts a string as its input:

```go
func HandlerFunc(s string) {
   fmt.Printf("Received: %s\n", s) 
}
```

Finally, set worker.WorkerFunc to your custom function from above and tell worker to process the string array:

```go
worker.WorkerFunc = HandlerFunc
    
var lines = []string{"foo", "bar", "car"}

worker.ProcessStringArray(lines)
```

Here is a full example:

```go
package main

import (
    "fmt"
    "github.com/clarke/golang-worker"
)

func main() {
    worker.WorkerFunc = HandlerFunc
    
    var lines = []string{"foo", "bar", "car"}
    
    worker.ProcessStringArray(lines)
}

func HandlerFunc(s string) {
   fmt.Printf("Received: %s\n", s)
}
```
