package tests

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"

	gongsim_fullstack "github.com/fullstack-lang/gongsim/go/fullstack"
	gongsim_static "github.com/fullstack-lang/gongsim/go/static"
)

var LadybugRadius = 0.00005 // a ladybug is 1mm wide

var maxDistanceInOneStep = float64(0.99 * (LadybugRadius / 4.0) * 60.0)

// cannot move from a half radius
var simulationStep = time.Microsecond * time.Duration(maxDistanceInOneStep*1000000.0)

var AbsoluteSpeed = 1.0 / 60.0 // a ladybug is 1m par minute

type LadybugSimulation struct {
}

var nbOfCollision = 0

func (specificEngine *LadybugSimulation) EventFired(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {
	return true
}
func (specificEngine *LadybugSimulation) Reset(engine *gongsim_models.Engine)          {}
func (specificEngine *LadybugSimulation) CommitAgents(engine *gongsim_models.Engine)   {}
func (specificEngine *LadybugSimulation) CheckoutAgents(engine *gongsim_models.Engine) {}
func (specificEngine *LadybugSimulation) GetLastCommitNb() uint                        { return 0 }
func (specificEngine *LadybugSimulation) GetLastCommitNbFromFront() uint               { return 0 }

type Ladybug struct {
	// Agent contains fields necessary for the simulation
	// You can reuse this struct
	gongsim_models.Agent

	Name string // necessary because it is a gongstruct

	UniqueId int // usefull for sorting between lady bus

	Position float64 // between 0.0 and 1.0

	Speed float64 // either AbsoluteSpeed or -AbsoluteSpeed
}

var eventNb = 0

const nbLadybugs = 32

var ladybugs []*Ladybug

// FireNextEvent is the minimal function that is necessary to run the simulation
func (ladybug *Ladybug) FireNextEvent() {

	event, eventTime := ladybug.GetNextEventAndRemoveIt()
	// log.Println("Ladybug", "FireNextEvent", ladybug.UniqueId, eventTime.Format("15:04:05.000000"))
	if eventNb%5000 == 0 && ladybug.UniqueId == 0 {
		log.Printf("Event %10d Time : %s Ladybug %s Position %10f Speed %10f", eventNb, eventTime.Format("15:04:05.000000"), ladybug.Name, ladybug.Position, ladybug.Speed)
	}

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
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

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
			return
		}

		// check for colisions (and compute)
		for _, otherLadybug := range ladybugs {
			// sum speeds
			sumOfSpeeds = sumOfSpeeds + math.Abs(otherLadybug.Speed)

			// do not compute collision of a ladybug with itslef
			if otherLadybug.UniqueId == ladybug.UniqueId {
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

				if ladybug.UniqueId == 0 {
					log.Printf("Event %10d Time : %s Pos %10f dist %10f ladybug %2d / %2d",
						eventNb, eventTime.Format("15:04:05.000000"), otherLadybug.Position, deltaX, ladybug.UniqueId, otherLadybug.UniqueId)
				}

				if deltaX > 0 && ladybug.Speed > 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.UniqueId)*1.0
				}
				if deltaX < 0 && ladybug.Speed < 0 {
					// return
					ladybug.Speed = -ladybug.Speed
					ladybug.Position = 10.0 + float64(ladybug.UniqueId)*1.0
				}

				nbOfCollision = nbOfCollision + 1
			}
		}

		if ladybug.Position < 0 || ladybug.Position > 1.0 {
			ladybug.Speed = 0.0
		}

		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		ladybug.QueueEvent(checkStateEvent)
	}
}

func TestLadybugSim(t *testing.T) {

	log.SetFlags(0)

	// rand.Seed(time.Now().UnixNano())

	// setup the static file server and get the controller
	r := gongsim_static.ServeStaticFiles(false)

	// setup stack
	gongsim_stage, _ := gongsim_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gongsim/go/models")

	engine := new(gongsim_models.Engine)
	engine.Name = "Simulation Engine"
	engine.Stage(gongsim_stage)
	gongsim_models.NewCommand(gongsim_stage, engine)

	// Kills Engine Simulation goroutine
	gongsim_models.Quit <- true

	// seven days of simulation
	engine.SetStartTime(time.Date(2021, time.July, 1, 0, 0, 0, 0, time.UTC))
	engine.SetCurrentTime(engine.GetStartTime())
	engine.State = gongsim_models.PAUSED
	engine.Speed = 1.0 // realtime
	// log.Printf("Sim start \t\t\t%s\n", engine.GetStartTime())

	// Three years
	engine.SetEndTime(time.Date(2021, time.July, 1, 0, 30, 0, 0, time.UTC))
	// log.Printf("Sim end  \t\t\t%s\n", engine.GetEndTime())

	// // PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	// var ladyBugSimulation LadybugSimulation
	// engine.Simulation = &ladyBugSimulation

	// initial positions of ladybugs cannot be close to each others than the radius
	initialXPosition := make(map[float64]*Ladybug)

	// append a ladybug agent to feed the discrete event engine
	for ladybugId := 0; ladybugId < nbLadybugs; ladybugId = ladybugId + 1 {
		ladyBug := new(Ladybug)
		ladyBug.Name = fmt.Sprintf("Ladybug #%2d", ladybugId)
		ladyBug.UniqueId = ladybugId

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

		engine.AppendAgent(ladyBug)
		var step gongsim_models.UpdateState
		step.SetFireTime(engine.GetStartTime())
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

	_, nextSimTime, _ := engine.GetNextEvent()
	for nextSimTime.Before(engine.GetEndTime()) {
		_, nextSimTime, _ = engine.FireNextEvent()
	}

	log.Printf("Ladybug sim over")
}
