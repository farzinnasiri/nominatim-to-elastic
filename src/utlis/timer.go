package utlis

import (
	"fmt"
	"sync"
	"time"
)

type timerType struct {
	Count     int
	TimeStamp time.Time
	Lock      sync.Mutex
}

var Timer = timerType{Count: 0}

func UpdateTimer(len int) {
	Timer.Lock.Lock()
	Timer.Count = Timer.Count + len
	Timer.TimeStamp = time.Now()
	Timer.Lock.Unlock()
}

var StartTime = time.Now()

func CalcTime(len int) {
	for {
		Timer.Lock.Lock()
		programTimer1 := Timer.Count
		Timer.Lock.Unlock()
		time.Sleep(time.Second)
		Timer.Lock.Lock()
		programTimer2 := Timer
		Timer.Lock.Unlock()
		speed := programTimer2.Count - programTimer1
		if speed != 0 {
			spendTime := programTimer2.TimeStamp.Sub(StartTime)
			leftCount := len - programTimer2.Count
			fmt.Printf("\r speed : %d per second,left time : %f second left.", speed,
				float64(leftCount)*spendTime.Seconds()/float64(programTimer2.Count))
		}

	}
}
