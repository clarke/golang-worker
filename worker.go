package worker

import (
    "fmt"
    "os"
    "runtime"
)

var WorkerFunc = stringNoop

func stringNoop(l string) {
    // do nothing
    fmt.Println("Set worker.WorkerFunc to something more useful.")
}

func ProcessStringArray(lines []string) {
    if WorkerFunc == nil {
        fmt.Println("worker.WorkerFunc not defined. Exiting.")
        os.Exit(1)
    }
    queue := make(chan *string)

    ncpu := runtime.NumCPU()
    if len(lines) < ncpu {
        ncpu = len(lines)
    }
    runtime.GOMAXPROCS(ncpu)

    // spawn workers
    for i := 0; i < ncpu; i++ {
        go Worker(i, queue)
    }

    // master: give work
    for i, _ := range lines {
        cnt := (i - ncpu + 1)
        queue <- &lines[i]
    }

    // all work is done
    // signal workers there is no more work
    for n := 0; n < ncpu; n++ {
        queue <- nil
    }
}

func Worker(id int, queue chan *string) {
    var line *string
    for {
        // get work item (pointer) from the queue
        line = <-queue
        if line == nil {
            break
        }

        WorkerFunc(*line)
    }
}
