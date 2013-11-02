package worker

import (
    "testing"
)

var resultArray = make([]string, 0)

func AddOneToTheEndOfString(l string) {
    resultArray = append(resultArray, l)
}

func TestProcessStringArray(t *testing.T) {
    lines := []string{"one", "two", "three"}

    if len(resultArray) != 0 {
        t.Error("Not starting with a clean resultArray")
    }

    WorkerFunc = AddOneToTheEndOfString
    ProcessStringArray(lines)

    if len(resultArray) != len(lines) {
        t.Error("Failed to populate resultArray with lines", resultArray)
    }
}

func TestDefaultWokerFunc(t *testing.T) {
    lines := []string{"one", "two", "three"}

    // Make sure to reset results so we can test to make sure it didn't do anything
    resultArray = []string{}

    // Set WorkerFunc back to the default
    WorkerFunc = stringNoop
    ProcessStringArray(lines)

    if len(resultArray) > 0 {
        t.Error("ProcessStringArray populated when it wasn't supposed to", resultArray)
    }
}
