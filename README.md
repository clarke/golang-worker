# Summary

Golang package for processing items using a worker queue model. The main point is to only spawn as many workers as there are CPUs available. The tasks to be completed will be added to a queue and as each worker completes a task, it checks the queue for remaining items.

This was *heavily* influenced by:

[Golang: Master/Worker in Go](http://devcry.heiho.net/2012/07/golang-masterworker-in-go.html)

My main goal was to make it as reusable as possible.

# Todo

* Tests!
