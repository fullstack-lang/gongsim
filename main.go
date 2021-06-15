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

	"github.com/fullstack-lang/gongsim/go/controllers"
	"github.com/fullstack-lang/gongsim/go/models"
	"github.com/fullstack-lang/gongsim/go/orm"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	play         = flag.Bool("play", false, "start rigth away")
	displayWatch = flag.Bool("displayWatch", false, "if true, print current status every 1/2 seconds")

	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
)

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

	// setup GORM
	db := orm.SetupModels(*logDBFlag, ":memory:")
	dbDB, err := db.DB()

	// since gongsim is a multi threaded application. It is important to set up
	// only one open connexion at a time
	if err != nil {
		panic("cannot access DB of db" + err.Error())
	}
	dbDB.SetMaxOpenConns(1)
	orm.BackRepo.Init(db)

	controllers.RegisterControllers(r)
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
	models.Stage.Commit()

	// provide the static route for the angular pages
	r.Use(static.Serve("/", EmbedFolder(ng, "ng/dist/ng")))
	r.NoRoute(func(c *gin.Context) {
		fmt.Println(c.Request.URL.Path, "doesn't exists, redirect on /")
		c.Redirect(http.StatusMovedPermanently, "/")
		c.Abort()
	})

	log.Printf("Server ready serve on localhost:8080")

	r.Run()
}

//go:embed ng/dist/ng
var ng embed.FS

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
