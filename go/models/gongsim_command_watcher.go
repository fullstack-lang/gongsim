package models

import (
	"log"
	"time"
)

var lastSimTime = EngineSingloton.currentTime

var DisplayWatch bool

// the watcher thread inspects the status of the simulation
func (gongsimCommand *GongsimCommand) watcher() {

	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	//
	// checkoutSchedulerPeriod is the period of the "checkout scheduler"
	watcherPeriod := 500 * time.Millisecond
	var WatcherSchedulerPeriod = time.NewTicker(watcherPeriod)
	for {
		select {
		case t := <-WatcherSchedulerPeriod.C:

			_ = t

			const layout = "Jan 2, 2006 at 15:04:05 (MST)"
			measuredSimSpeed := float64(EngineSingloton.currentTime.Sub(lastSimTime)) / float64(watcherPeriod)
			if DisplayWatch {
				log.Printf("simulated time is %s, status %s, speed %f, speed request %f",
					EngineSingloton.currentTime.Format(layout), EngineSingloton.State, measuredSimSpeed, EngineSingloton.Speed)
			}
			lastSimTime = EngineSingloton.currentTime
		}
	}
}
