package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongsim/go/fullstack"
	"github.com/fullstack-lang/gongsim/go/models"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"

	gongsim "github.com/fullstack-lang/gongsim"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnCommit = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

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

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
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

	if *cpuprofile != "" {
		models.CpuProfile = true
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// setup controlers
	if !*logGINFlag {
		myfile, _ := os.Create("/tmp/server.log")
		gin.DefaultWriter = myfile
	}
	r := gin.Default()
	r.Use(cors.Default())

	// setup stack
	var stage *models.StageStruct
	if *marshallOnCommit != "" {
		stage, _ = fullstack.NewStackInstance(r, "")
	} else {
		stage, _ = fullstack.NewStackInstance(r, "", "")
	}

	// hook automatic marshall to go code at every commit
	if *marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		stage.OnInitCommitFromFrontCallback = hook
	}

	gongdoc_load.Load(
		"gongsim",
		"github.com/fullstack-lang/gongsim/go/models",
		gongsim.GoDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

	models.EngineSingloton.ControlMode = models.CLIENT_CONTROL

	// seven days of simulation
	models.EngineSingloton.SetStartTime(time.Date(1676, time.January, 1, 0, 0, 0, 0, time.UTC))
	models.EngineSingloton.SetCurrentTime(models.EngineSingloton.GetStartTime())
	models.EngineSingloton.State = models.PAUSED
	models.EngineSingloton.Speed = 0.5 * 24 * 3600.0 // days per second
	log.Printf("Sim start \t\t\t%s\n", models.EngineSingloton.GetStartTime())

	// Three years
	models.EngineSingloton.SetEndTime(time.Date(1680, time.January, 1, 0, 0, 0, 0, time.UTC))
	log.Printf("Sim end  \t\t\t%s\n", models.EngineSingloton.GetEndTime())

	// PLUMBING n°1: callback for treating model specific action. In this case, see specific engine
	var simulation models.Simulation
	models.EngineSingloton.Simulation = &simulation

	// append a dummy agent to feed the discrete event engine with at least an event
	dummyAgent := new(models.DummyAgent)
	dummyAgent.Name = "Dummy"

	models.EngineSingloton.AppendAgent(dummyAgent)
	var step models.UpdateState
	step.SetFireTime(models.EngineSingloton.GetStartTime())
	step.Period = 1 * time.Second //
	step.Name = "update of planetary motion"
	dummyAgent.QueueEvent(&step)

	// start right away
	if *play {
		models.GongsimCommandSingloton.Command = models.COMMAND_PLAY
	}
	if *displayWatch {
		models.DisplayWatch = true
	}

	// commit simulation stage
	stage.Commit()

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(gongsim.NgDistNg, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
