package asyncservice

import (
	"testing"
	"math"
        "github.com/dnikiforov1967/sttest/config"
        "github.com/stretchr/testify/assert"
)

//Test round() implementation
func TestRound(t * testing.T) {
	array := []float64{1.005, 1.005001, 1.0049999}

	for _, value := range array {
		res := Round2(value)
		if res != math.Round(value*100)/100 {
			t.Errorf("Incorrect rounding of %f, result is %f",value,res)
		}
	}
}

func TestTimeout(t * testing.T) {
    config.TimeOut = 2000
    sigChan := make(chan int)
    go proceed(1,"A",0.01,0.001,sigChan)
    <-sigChan
    task, _ := getTaskState(1)
    assert.Equal(t, task.Status, StatusTimedOut, "StatusTimedOut is expected")
}

func TestNormalRun(t * testing.T) {
    config.TimeOut = 6000
    sigChan := make(chan int)
    go proceed(2,"A",0.01,0.001,sigChan)
    <-sigChan
    task, _ := getTaskState(2)
    assert.Equal(t, task.Status, StatusCompleted, "StatusCompleted is expected")
}