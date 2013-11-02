package worker

import (
    "testing"
)

func TestProcessStringArray(t *testing.T) {
    // Save the default WorkerFunc so that we can reset it after the test
    var defaultWorkerFunc = WorkerFunc

    lines := []string{"one", "two", "three"}

    var resultArray = make([]string, 0)
    WorkerFunc = func (l string) {
        resultArray = append(resultArray, l)
    }

    if len(resultArray) != 0 {
        t.Error("Not starting with a clean resultArray")
    }

    ProcessStringArray(lines)

    if len(resultArray) != len(lines) {
        t.Error("Failed to populate resultArray with lines", resultArray)
    }

    // Set WorkerFunc back to the default
    WorkerFunc = defaultWorkerFunc
}
