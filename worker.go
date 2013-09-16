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
        // Print "100" at the completion of about every 100 items.
        // Since we are threading, it's difficult to tell exactly
        // when the 100th item has been processed, so this is
        // more of a guestimated progress meter than an actual
        // mile marker.
        cnt := (i - ncpu + 1)
        if cnt > 0 && cnt%100 == 0 {
            fmt.Printf("%d", cnt)
        }
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
