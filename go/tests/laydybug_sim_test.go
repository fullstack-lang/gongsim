package tests

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	"github.com/fullstack-lang/gongsim/go/models"
)

var LadybugRadius = 0.00005 // a ladybug is 1mm wide

var maxDistanceInOneStep = float64(0.99 * (LadybugRadius / 4.0) * 60.0)

// cannot move from a half radius
var simulationStep = time.Microsecond * time.Duration(maxDistanceInOneStep*1000000.0)

var AbsoluteSpeed = 1.0 / 60.0 // a ladybug is 1m par minute

type LadybugSimulation struct {
}

var nbOfCollision = 0

func (specificEngine *LadybugSimulation) EventFired(engine *models.Engine)              {}
func (specificEngine *LadybugSimulation) HasAnyStateChanged(engine *models.Engine) bool { return true }
func (specificEngine *LadybugSimulation) Reset(engine *models.Engine)                   {}
func (specificEngine *LadybugSimulation) CommitAgents(engine *models.Engine)            {}
func (specificEngine *LadybugSimulation) CheckoutAgents(engine *models.Engine)          {}
func (specificEngine *LadybugSimulation) GetLastCommitNb() uint                         { return 0 }
func (specificEngine *LadybugSimulation) GetLastCommitNbFromFront() uint                { return 0 }

type Ladybug struct {
	models.Agent

	Name string

	Id int

	Position float64 // between 0.0 and 1.0

	Speed float64 // either AbsoluteSpeed or -AbsoluteSpeed
}

var eventNb = 0

const nbLadybugs = 32

var ladybugs []*Ladybug

func (ladybug *Ladybug) FireNextEvent() {
	event, eventTime := ladybug.GetNextEventAndRemoveIt()

	if eventNb%32 == 0 && ladybug.Speed != 0.0 {

		var positions string
		for _, _ladybug := range ladybugs {
			direction := "D"
			if _ladybug.Speed < 0 {
				direction = "G"
			}
			positions = positions + fmt.Sprintf("%0.2f %s", _ladybug.Position, direction)
		}

		// log.Printf("%s", positions)
	}

	switch event.(type) {
	case *models.UpdateState:
		checkStateEvent := event.(*models.UpdateState)

		// if eventNb%5000 == 0 && ladybug.Id == 0 {
		// 	log.Printf("Event %10d Time : %s Ladybug %s Position %10f Speed %10f", eventNb, eventTime.Format("15:04:05.000000"), ladybug.Name, ladybug.Position, ladybug.Speed)
		// }
		// if eventNb%3200 == 0 && ladybug.Speed != 0.0 {
		// 	log.Printf("Event %10d Time : %s Ladybug %s Position %10f Speed %10f", eventNb, eventTime.Format("15:04:05.000000"), ladybug.Name, ladybug.Position, ladybug.Speed)
		// }
		eventNb = eventNb + 1

		ladybug.Position = ladybug.Position + simulationStep.Seconds()*ladybug.Speed

		// stop simu if sum of speeds is 0
		sumOfSpeeds := 0.0
		for _, otherLadybug := range ladybugs {
			// sum speeds
			sumOfSpeeds = sumOfSpeeds + math.Abs(otherLadybug.Speed)
		}
		if sumOfSpeeds == 0 {
			log.Printf("Event %10d Time : %s, nbOfCollisions %d simulation over",
				eventNb, eventTime.Format("15:04:05.000000"), nbOfCollision/2)
			os.Exit(0)
		}

		// check for colisions (and compute)
		for _, otherLadybug := range ladybugs {
			// sum speeds
			sumOfSpeeds = sumOfSpeeds + math.Abs(otherLadybug.Speed)

			// do not compute collision of a ladybug with itslef
			if otherLadybug.Id == ladybug.Id {
				continue
			}

			// do not compute collision if the other ladybug is allready out
			if otherLadybug.Speed == 0.0 {
				continue
			}

			// get the between X positions
			deltaX := otherLadybug.Position - ladybug.Position

			// there is a collision if both are within a Ladybug diameter
			if math.Abs(deltaX) < 2*LadybugRadius {

				if ladybug.Id == 0 {
					log.Printf("Event %10d Time : %s Pos %10f dist %10f ladybug %2d / %2d",
						eventNb, eventTime.Format("15:04:05.000000"), otherLadybug.Position, deltaX, ladybug.Id, otherLadybug.Id)
				}

				if deltaX > 0 && ladybug.Speed > 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.Id)*1.0
				}
				if deltaX < 0 && ladybug.Speed < 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.Id)*1.0
				}

				nbOfCollision = nbOfCollision + 1
			}
		}

		if ladybug.Position < 0 || ladybug.Position > 1.0 {
			ladybug.Speed = 0.0
		}

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		ladybug.QueueEvent(checkStateEvent)
	}
}

func TestLadybugSim(t *testing.T) {

	log.SetFlags(0)

	// rand.Seed(time.Now().UnixNano())

	// Kills Engine Simulation goroutine
	models.Quit <- true

	// seven days of simulation
	models.EngineSingloton.SetStartTime(time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC))
	models.EngineSingloton.SetCurrentTime(models.EngineSingloton.GetStartTime())
	models.EngineSingloton.State = models.PAUSED
	models.EngineSingloton.Speed = 1.0 // realtime
	// log.Printf("Sim start \t\t\t%s\n", models.EngineSingloton.GetStartTime())

	// Three years
	models.EngineSingloton.SetEndTime(time.Date(2021, time.July, 1, 0, 30, 0, 0, time.UTC))
	// log.Printf("Sim end  \t\t\t%s\n", models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var ladyBugSimulation LadybugSimulation
	models.EngineSingloton.Simulation = &ladyBugSimulation

	// initial positions of ladybugs cannot be close to each others than the radius
	initialXPosition := make(map[float64]*Ladybug)

	// append a ladybug agent to feed the discrete event engine
	for ladybugId := 0; ladybugId < nbLadybugs; ladybugId = ladybugId + 1 {
		ladyBug := new(Ladybug)
		ladyBug.Name = fmt.Sprintf("Ladybug #%2d", ladybugId)
		ladyBug.Id = ladybugId

		ladybugs = append(ladybugs, ladyBug)

		// set up position
		positionX := rand.Float64()

		// adjust it on a multiple of the ladybug diameter
		positionX = math.Round(positionX*1.0/LadybugRadius) * LadybugRadius

		ladyBug.Position = positionX
		if initialXPosition[positionX] != nil {
			log.Panic("same initial position")
		}

		// decide orientaiton of the speed
		if rand.Float64() > 0.5 {
			ladyBug.Speed = AbsoluteSpeed
		} else {
			ladyBug.Speed = -AbsoluteSpeed
		}

		models.EngineSingloton.AppendAgent(ladyBug)
		var step models.UpdateState
		step.SetFireTime(models.EngineSingloton.GetStartTime())
		step.Period = simulationStep //
		step.Name = "update of laybug motion"
		ladyBug.QueueEvent(&step)
	}

	sort.Slice(ladybugs[:], func(i, j int) bool {
		return ladybugs[i].Position < ladybugs[j].Position
	})

	// compute left & right relay positions
	leftRelayInitialPosX := 0.0
	for _, ladybug := range ladybugs {
		if ladybug.Speed > 0 {
			leftRelayInitialPosX = ladybug.Position
			log.Printf("Left relay %1.4f", leftRelayInitialPosX)
			break
		}
	}

	rightRelayInitialPosX := 0.0
	for i := len(ladybugs) - 1; i >= 0; i-- {

		ladybug := ladybugs[i]
		if ladybug.Speed > 0 {
			rightRelayInitialPosX = ladybug.Position
			log.Printf("right relay %1.4f", rightRelayInitialPosX)
			break
		}
	}

	middleBetweenRelayX := (leftRelayInitialPosX + rightRelayInitialPosX) / 2.0
	log.Printf("middle between relays %1.4f", middleBetweenRelayX)

	// time for the left relay to get out
	leftRelayDistanceToExit := (middleBetweenRelayX - leftRelayInitialPosX) + middleBetweenRelayX
	log.Printf("left distance to exit %1.4f", leftRelayDistanceToExit)
	log.Printf("left time to exit %1.4f", leftRelayDistanceToExit*60.0)

	rightRelayDistanceToExit := (rightRelayInitialPosX - middleBetweenRelayX) + (1.0 - middleBetweenRelayX)
	log.Printf("right distance to exit %1.4f", rightRelayDistanceToExit)
	log.Printf("right time to exit %1.4f", rightRelayDistanceToExit*60.0)

	maxRelayDistanceFromBorder := math.Max(leftRelayInitialPosX, 1.0-rightRelayInitialPosX)
	log.Printf("max distance relay to border %1.4f", maxRelayDistanceFromBorder)

	_, nextSimTime, _ := models.EngineSingloton.GetNextEvent()
	for nextSimTime.Before(models.EngineSingloton.GetEndTime()) {
		_, nextSimTime, _ = models.EngineSingloton.FireNextEvent()
	}

	log.Printf("Ladybug sim over")
}
