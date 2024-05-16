package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	gongsim_go "github.com/fullstack-lang/gongsim/go"
	gongsim_fullstack "github.com/fullstack-lang/gongsim/go/fullstack"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
	gongsim_probe "github.com/fullstack-lang/gongsim/go/probe"
	gongsim_static "github.com/fullstack-lang/gongsim/go/static"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup  = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	unmarshall         = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	play         = flag.Bool("play", false, "start rigth away")
	displayWatch = flag.Bool("displayWatch", false, "if true, print current status every 1/2 seconds")

	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gongsim_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gongsim/go/models", "main")
}

func main() {

	log.SetPrefix("gongsim: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gongsim_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	stage, gongsimBackRepo := gongsim_fullstack.NewStackInstance(
		r, "gongsim")

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Checkout()
		stage.Marshall(file, "github.com/fullstack-lang/gongsim/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}
		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	if *unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := gongsim_models.ParseAstFile(stage, *unmarshallFromCode)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		stage.OnInitCommitFromFrontCallback = hook
	}

	engine := new(gongsim_models.Engine)
	engine.Name = "Simulation Engine"
	engine.DisplayFormat = "Jan 2006"
	engine.Stage(stage)

	// the gongsim command orchestrates the simulation engine regarding to the
	// the rest of the stack. It manages when the stage has to be commited to the
	// back repo or when the back repo has to be checked out to the stage
	gongsimCommand := gongsim_models.NewGongSimCommand(stage, engine)
	_ = gongsimCommand

	// seven days of simulation
	engine.SetStartTime(time.Date(1676, time.January, 1, 0, 0, 0, 0, time.UTC))
	engine.SetCurrentTime(engine.GetStartTime())
	engine.State = gongsim_models.PAUSED
	engine.Speed = 0.5 * 24 * 3600.0 // days per second
	log.Printf("Sim start \t\t\t%s\n", engine.GetStartTime())

	// Three years
	engine.SetEndTime(time.Date(1680, time.January, 1, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", engine.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var simulation gongsim_models.Simulation
	engine.Simulation = &simulation

	// append a dummy agent to feed the discrete event engine with at least an event
	dummyAgent := new(gongsim_models.DummyAgent)
	dummyAgent.Name = "Dummy"

	engine.AppendAgent(dummyAgent)
	var step gongsim_models.UpdateState
	step.SetFireTime(engine.GetStartTime())
	step.Period = 1 * time.Second //
	step.Name = "update of planetary motion"
	dummyAgent.QueueEvent(&step)

	// start right away
	if *play {
		gongsimCommand.Command = gongsim_models.COMMAND_PLAY
	}
	if *displayWatch {
		gongsim_models.DisplayWatch = true
	}

	// commit simulation stage
	stage.Commit()

	gongsim_probe.NewProbe(r, gongsim_go.GoModelsDir, gongsim_go.GoDiagramsDir,
		*embeddedDiagrams, "gongsim", stage, gongsimBackRepo)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
