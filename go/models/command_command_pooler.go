package models

import (
	"log"
	"time"
)

// commandPooler" is a simple scheduled thread that manages the checkout
// of the current command from the back repo to the stage. If the front client writes a "PLAY" or "PAUSE"
// command, it is written to the back repo and the command pooler retrieves this command to the stage where it
// can be read by the "engine driver". The "command pooler" is a scheduled task
func (command *Command) commandPooler() {

	// commandPoolerPeriod is the period of the "command pooler"
	//
	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	var CommandPoolerPeriod = time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case t := <-CommandPoolerPeriod.C:

			command.Checkout(command.stage)
			if command.gongsimStatus.CompletionDate != command.CommandDate {
				log.Println("commandPooler reads new command ", command.Command, "  timestamp ", command.CommandDate, " at ", t)

				command.gongsimStatus.CurrentCommand = command.Command
				command.gongsimStatus.CompletionDate = command.CommandDate
				command.stage.Commit()
			}
			if command.gongsimStatus.SpeedCommandCompletionDate != command.DateSpeedCommand {
				log.Println("commandPooler reads new speed command ", command.SpeedCommandType, "  timestamp ", command.CommandDate, " at ", t)

				switch command.SpeedCommandType {
				case COMMAND_DECREASE_SPEED_50_PERCENTS:
					command.Engine.Speed *= 0.5
					command.stage.Commit()
				case COMMAND_INCREASE_SPEED_100_PERCENTS:
					command.Engine.Speed *= 2.0
					command.stage.Commit()
				}

				command.gongsimStatus.CurrentSpeedCommand = command.SpeedCommandType
				command.gongsimStatus.SpeedCommandCompletionDate = command.DateSpeedCommand
				command.stage.Commit()
			}
		}
	}
}
