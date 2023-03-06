package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongsim/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_NewDiagram models.StageStruct
var ___dummy__Time_NewDiagram time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_NewDiagram ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_NewDiagram map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.ControlMode": ref_models.ControlMode(""),

	"ref_models.Engine": &(ref_models.Engine{}),

	"ref_models.Engine.ControlMode": (ref_models.Engine{}).ControlMode,

	"ref_models.Engine.CurrentTime": (ref_models.Engine{}).CurrentTime,

	"ref_models.Engine.EndTime": (ref_models.Engine{}).EndTime,

	"ref_models.Engine.Fired": (ref_models.Engine{}).Fired,

	"ref_models.Engine.Name": (ref_models.Engine{}).Name,

	"ref_models.Engine.SecondsSinceStart": (ref_models.Engine{}).SecondsSinceStart,

	"ref_models.Engine.Speed": (ref_models.Engine{}).Speed,

	"ref_models.Engine.State": (ref_models.Engine{}).State,

	"ref_models.EngineRunMode": ref_models.EngineRunMode(0),

	"ref_models.EngineState": ref_models.EngineState(""),

	"ref_models.Event": &(ref_models.Event{}),

	"ref_models.Event.Duration": (ref_models.Event{}).Duration,

	"ref_models.Event.Name": (ref_models.Event{}).Name,

	"ref_models.GongsimCommand": &(ref_models.GongsimCommand{}),

	"ref_models.GongsimCommand.Command": (ref_models.GongsimCommand{}).Command,

	"ref_models.GongsimCommand.CommandDate": (ref_models.GongsimCommand{}).CommandDate,

	"ref_models.GongsimCommand.DateSpeedCommand": (ref_models.GongsimCommand{}).DateSpeedCommand,

	"ref_models.GongsimCommand.Name": (ref_models.GongsimCommand{}).Name,

	"ref_models.GongsimCommand.SpeedCommandType": (ref_models.GongsimCommand{}).SpeedCommandType,

	"ref_models.GongsimStatus": &(ref_models.GongsimStatus{}),

	"ref_models.GongsimStatus.CompletionDate": (ref_models.GongsimStatus{}).CompletionDate,

	"ref_models.GongsimStatus.CurrentCommand": (ref_models.GongsimStatus{}).CurrentCommand,

	"ref_models.GongsimStatus.CurrentSpeedCommand": (ref_models.GongsimStatus{}).CurrentSpeedCommand,

	"ref_models.GongsimStatus.Name": (ref_models.GongsimStatus{}).Name,

	"ref_models.GongsimStatus.SpeedCommandCompletionDate": (ref_models.GongsimStatus{}).SpeedCommandCompletionDate,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["NewDiagram"] = NewDiagramInjection
// }

// NewDiagramInjection will stage objects of database "NewDiagram"
func NewDiagramInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_NewDiagram := (&models.Classdiagram{Name: `NewDiagram`}).Stage()

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_Command := (&models.Field{Name: `Command`}).Stage()
	__Field__000001_CommandDate := (&models.Field{Name: `CommandDate`}).Stage()
	__Field__000002_CompletionDate := (&models.Field{Name: `CompletionDate`}).Stage()
	__Field__000003_ControlMode := (&models.Field{Name: `ControlMode`}).Stage()
	__Field__000004_CurrentCommand := (&models.Field{Name: `CurrentCommand`}).Stage()
	__Field__000005_CurrentSpeedCommand := (&models.Field{Name: `CurrentSpeedCommand`}).Stage()
	__Field__000006_CurrentTime := (&models.Field{Name: `CurrentTime`}).Stage()
	__Field__000007_DateSpeedCommand := (&models.Field{Name: `DateSpeedCommand`}).Stage()
	__Field__000008_Duration := (&models.Field{Name: `Duration`}).Stage()
	__Field__000009_EndTime := (&models.Field{Name: `EndTime`}).Stage()
	__Field__000010_Fired := (&models.Field{Name: `Fired`}).Stage()
	__Field__000011_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000012_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000013_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000014_Name := (&models.Field{Name: `Name`}).Stage()
	__Field__000015_SecondsSinceStart := (&models.Field{Name: `SecondsSinceStart`}).Stage()
	__Field__000016_Speed := (&models.Field{Name: `Speed`}).Stage()
	__Field__000017_SpeedCommandCompletionDate := (&models.Field{Name: `SpeedCommandCompletionDate`}).Stage()
	__Field__000018_SpeedCommandType := (&models.Field{Name: `SpeedCommandType`}).Stage()
	__Field__000019_State := (&models.Field{Name: `State`}).Stage()

	// Declarations of staged instances of GongEnumShape
	__GongEnumShape__000000_NewDiagram_ControlMode := (&models.GongEnumShape{Name: `NewDiagram-ControlMode`}).Stage()
	__GongEnumShape__000001_NewDiagram_EngineRunMode := (&models.GongEnumShape{Name: `NewDiagram-EngineRunMode`}).Stage()
	__GongEnumShape__000002_NewDiagram_EngineState := (&models.GongEnumShape{Name: `NewDiagram-EngineState`}).Stage()

	// Declarations of staged instances of GongEnumValueEntry
	__GongEnumValueEntry__000000_AUTONOMOUS := (&models.GongEnumValueEntry{Name: `AUTONOMOUS`}).Stage()
	__GongEnumValueEntry__000001_CLIENT_CONTROL := (&models.GongEnumValueEntry{Name: `CLIENT_CONTROL`}).Stage()
	__GongEnumValueEntry__000002_OVER := (&models.GongEnumValueEntry{Name: `OVER`}).Stage()
	__GongEnumValueEntry__000003_PAUSED := (&models.GongEnumValueEntry{Name: `PAUSED`}).Stage()
	__GongEnumValueEntry__000004_RUNNING := (&models.GongEnumValueEntry{Name: `RUNNING`}).Stage()

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_NewDiagram_Engine := (&models.GongStructShape{Name: `NewDiagram-Engine`}).Stage()
	__GongStructShape__000001_NewDiagram_Event := (&models.GongStructShape{Name: `NewDiagram-Event`}).Stage()
	__GongStructShape__000002_NewDiagram_GongsimCommand := (&models.GongStructShape{Name: `NewDiagram-GongsimCommand`}).Stage()
	__GongStructShape__000003_NewDiagram_GongsimStatus := (&models.GongStructShape{Name: `NewDiagram-GongsimStatus`}).Stage()

	// Declarations of staged instances of Link

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_NewDiagram_ControlMode := (&models.Position{Name: `Pos-NewDiagram-ControlMode`}).Stage()
	__Position__000001_Pos_NewDiagram_Engine := (&models.Position{Name: `Pos-NewDiagram-Engine`}).Stage()
	__Position__000002_Pos_NewDiagram_EngineRunMode := (&models.Position{Name: `Pos-NewDiagram-EngineRunMode`}).Stage()
	__Position__000003_Pos_NewDiagram_EngineState := (&models.Position{Name: `Pos-NewDiagram-EngineState`}).Stage()
	__Position__000004_Pos_NewDiagram_Event := (&models.Position{Name: `Pos-NewDiagram-Event`}).Stage()
	__Position__000005_Pos_NewDiagram_GongsimCommand := (&models.Position{Name: `Pos-NewDiagram-GongsimCommand`}).Stage()
	__Position__000006_Pos_NewDiagram_GongsimStatus := (&models.Position{Name: `Pos-NewDiagram-GongsimStatus`}).Stage()

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_NewDiagram.Name = `NewDiagram`
	__Classdiagram__000000_NewDiagram.IsInDrawMode = true

	// Field values setup
	__Field__000000_Command.Name = `Command`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.Command]
	__Field__000000_Command.Identifier = `ref_models.GongsimCommand.Command`
	__Field__000000_Command.FieldTypeAsString = ``
	__Field__000000_Command.Structname = `GongsimCommand`
	__Field__000000_Command.Fieldtypename = `GongsimCommandType`

	// Field values setup
	__Field__000001_CommandDate.Name = `CommandDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.CommandDate]
	__Field__000001_CommandDate.Identifier = `ref_models.GongsimCommand.CommandDate`
	__Field__000001_CommandDate.FieldTypeAsString = ``
	__Field__000001_CommandDate.Structname = `GongsimCommand`
	__Field__000001_CommandDate.Fieldtypename = `string`

	// Field values setup
	__Field__000002_CompletionDate.Name = `CompletionDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CompletionDate]
	__Field__000002_CompletionDate.Identifier = `ref_models.GongsimStatus.CompletionDate`
	__Field__000002_CompletionDate.FieldTypeAsString = ``
	__Field__000002_CompletionDate.Structname = `GongsimStatus`
	__Field__000002_CompletionDate.Fieldtypename = `string`

	// Field values setup
	__Field__000003_ControlMode.Name = `ControlMode`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.ControlMode]
	__Field__000003_ControlMode.Identifier = `ref_models.Engine.ControlMode`
	__Field__000003_ControlMode.FieldTypeAsString = ``
	__Field__000003_ControlMode.Structname = `Engine`
	__Field__000003_ControlMode.Fieldtypename = `ControlMode`

	// Field values setup
	__Field__000004_CurrentCommand.Name = `CurrentCommand`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CurrentCommand]
	__Field__000004_CurrentCommand.Identifier = `ref_models.GongsimStatus.CurrentCommand`
	__Field__000004_CurrentCommand.FieldTypeAsString = ``
	__Field__000004_CurrentCommand.Structname = `GongsimStatus`
	__Field__000004_CurrentCommand.Fieldtypename = `GongsimCommandType`

	// Field values setup
	__Field__000005_CurrentSpeedCommand.Name = `CurrentSpeedCommand`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.CurrentSpeedCommand]
	__Field__000005_CurrentSpeedCommand.Identifier = `ref_models.GongsimStatus.CurrentSpeedCommand`
	__Field__000005_CurrentSpeedCommand.FieldTypeAsString = ``
	__Field__000005_CurrentSpeedCommand.Structname = `GongsimStatus`
	__Field__000005_CurrentSpeedCommand.Fieldtypename = `SpeedCommandType`

	// Field values setup
	__Field__000006_CurrentTime.Name = `CurrentTime`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.CurrentTime]
	__Field__000006_CurrentTime.Identifier = `ref_models.Engine.CurrentTime`
	__Field__000006_CurrentTime.FieldTypeAsString = ``
	__Field__000006_CurrentTime.Structname = `Engine`
	__Field__000006_CurrentTime.Fieldtypename = `string`

	// Field values setup
	__Field__000007_DateSpeedCommand.Name = `DateSpeedCommand`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.DateSpeedCommand]
	__Field__000007_DateSpeedCommand.Identifier = `ref_models.GongsimCommand.DateSpeedCommand`
	__Field__000007_DateSpeedCommand.FieldTypeAsString = ``
	__Field__000007_DateSpeedCommand.Structname = `GongsimCommand`
	__Field__000007_DateSpeedCommand.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Duration.Name = `Duration`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event.Duration]
	__Field__000008_Duration.Identifier = `ref_models.Event.Duration`
	__Field__000008_Duration.FieldTypeAsString = ``
	__Field__000008_Duration.Structname = `Event`
	__Field__000008_Duration.Fieldtypename = `Duration`

	// Field values setup
	__Field__000009_EndTime.Name = `EndTime`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.EndTime]
	__Field__000009_EndTime.Identifier = `ref_models.Engine.EndTime`
	__Field__000009_EndTime.FieldTypeAsString = ``
	__Field__000009_EndTime.Structname = `Engine`
	__Field__000009_EndTime.Fieldtypename = `string`

	// Field values setup
	__Field__000010_Fired.Name = `Fired`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Fired]
	__Field__000010_Fired.Identifier = `ref_models.Engine.Fired`
	__Field__000010_Fired.FieldTypeAsString = ``
	__Field__000010_Fired.Structname = `Engine`
	__Field__000010_Fired.Fieldtypename = `int`

	// Field values setup
	__Field__000011_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.Name]
	__Field__000011_Name.Identifier = `ref_models.GongsimCommand.Name`
	__Field__000011_Name.FieldTypeAsString = ``
	__Field__000011_Name.Structname = `GongsimCommand`
	__Field__000011_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000012_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event.Name]
	__Field__000012_Name.Identifier = `ref_models.Event.Name`
	__Field__000012_Name.FieldTypeAsString = ``
	__Field__000012_Name.Structname = `Event`
	__Field__000012_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000013_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Name]
	__Field__000013_Name.Identifier = `ref_models.Engine.Name`
	__Field__000013_Name.FieldTypeAsString = ``
	__Field__000013_Name.Structname = `Engine`
	__Field__000013_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000014_Name.Name = `Name`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.Name]
	__Field__000014_Name.Identifier = `ref_models.GongsimStatus.Name`
	__Field__000014_Name.FieldTypeAsString = ``
	__Field__000014_Name.Structname = `GongsimStatus`
	__Field__000014_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000015_SecondsSinceStart.Name = `SecondsSinceStart`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.SecondsSinceStart]
	__Field__000015_SecondsSinceStart.Identifier = `ref_models.Engine.SecondsSinceStart`
	__Field__000015_SecondsSinceStart.FieldTypeAsString = ``
	__Field__000015_SecondsSinceStart.Structname = `Engine`
	__Field__000015_SecondsSinceStart.Fieldtypename = `float64`

	// Field values setup
	__Field__000016_Speed.Name = `Speed`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.Speed]
	__Field__000016_Speed.Identifier = `ref_models.Engine.Speed`
	__Field__000016_Speed.FieldTypeAsString = ``
	__Field__000016_Speed.Structname = `Engine`
	__Field__000016_Speed.Fieldtypename = `float64`

	// Field values setup
	__Field__000017_SpeedCommandCompletionDate.Name = `SpeedCommandCompletionDate`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus.SpeedCommandCompletionDate]
	__Field__000017_SpeedCommandCompletionDate.Identifier = `ref_models.GongsimStatus.SpeedCommandCompletionDate`
	__Field__000017_SpeedCommandCompletionDate.FieldTypeAsString = ``
	__Field__000017_SpeedCommandCompletionDate.Structname = `GongsimStatus`
	__Field__000017_SpeedCommandCompletionDate.Fieldtypename = `string`

	// Field values setup
	__Field__000018_SpeedCommandType.Name = `SpeedCommandType`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand.SpeedCommandType]
	__Field__000018_SpeedCommandType.Identifier = `ref_models.GongsimCommand.SpeedCommandType`
	__Field__000018_SpeedCommandType.FieldTypeAsString = ``
	__Field__000018_SpeedCommandType.Structname = `GongsimCommand`
	__Field__000018_SpeedCommandType.Fieldtypename = `SpeedCommandType`

	// Field values setup
	__Field__000019_State.Name = `State`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine.State]
	__Field__000019_State.Identifier = `ref_models.Engine.State`
	__Field__000019_State.FieldTypeAsString = ``
	__Field__000019_State.Structname = `Engine`
	__Field__000019_State.Fieldtypename = `EngineState`

	// GongEnumShape values setup
	__GongEnumShape__000000_NewDiagram_ControlMode.Name = `NewDiagram-ControlMode`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode]
	__GongEnumShape__000000_NewDiagram_ControlMode.Identifier = `ref_models.ControlMode`
	__GongEnumShape__000000_NewDiagram_ControlMode.Width = 240.000000
	__GongEnumShape__000000_NewDiagram_ControlMode.Heigth = 93.000000

	// GongEnumShape values setup
	__GongEnumShape__000001_NewDiagram_EngineRunMode.Name = `NewDiagram-EngineRunMode`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineRunMode]
	__GongEnumShape__000001_NewDiagram_EngineRunMode.Identifier = `ref_models.EngineRunMode`
	__GongEnumShape__000001_NewDiagram_EngineRunMode.Width = 240.000000
	__GongEnumShape__000001_NewDiagram_EngineRunMode.Heigth = 63.000000

	// GongEnumShape values setup
	__GongEnumShape__000002_NewDiagram_EngineState.Name = `NewDiagram-EngineState`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState]
	__GongEnumShape__000002_NewDiagram_EngineState.Identifier = `ref_models.EngineState`
	__GongEnumShape__000002_NewDiagram_EngineState.Width = 240.000000
	__GongEnumShape__000002_NewDiagram_EngineState.Heigth = 108.000000

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000000_AUTONOMOUS.Name = `AUTONOMOUS`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode.AUTONOMOUS]
	__GongEnumValueEntry__000000_AUTONOMOUS.Identifier = `ref_models.ControlMode.AUTONOMOUS`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000001_CLIENT_CONTROL.Name = `CLIENT_CONTROL`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.ControlMode.CLIENT_CONTROL]
	__GongEnumValueEntry__000001_CLIENT_CONTROL.Identifier = `ref_models.ControlMode.CLIENT_CONTROL`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000002_OVER.Name = `OVER`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.OVER]
	__GongEnumValueEntry__000002_OVER.Identifier = `ref_models.EngineState.OVER`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000003_PAUSED.Name = `PAUSED`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.PAUSED]
	__GongEnumValueEntry__000003_PAUSED.Identifier = `ref_models.EngineState.PAUSED`

	// GongEnumValueEntry values setup
	__GongEnumValueEntry__000004_RUNNING.Name = `RUNNING`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.EngineState.RUNNING]
	__GongEnumValueEntry__000004_RUNNING.Identifier = `ref_models.EngineState.RUNNING`

	// GongStructShape values setup
	__GongStructShape__000000_NewDiagram_Engine.Name = `NewDiagram-Engine`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Engine]
	__GongStructShape__000000_NewDiagram_Engine.Identifier = `ref_models.Engine`
	__GongStructShape__000000_NewDiagram_Engine.ShowNbInstances = false
	__GongStructShape__000000_NewDiagram_Engine.NbInstances = 0
	__GongStructShape__000000_NewDiagram_Engine.Width = 240.000000
	__GongStructShape__000000_NewDiagram_Engine.Heigth = 183.000000
	__GongStructShape__000000_NewDiagram_Engine.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_NewDiagram_Event.Name = `NewDiagram-Event`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Event]
	__GongStructShape__000001_NewDiagram_Event.Identifier = `ref_models.Event`
	__GongStructShape__000001_NewDiagram_Event.ShowNbInstances = false
	__GongStructShape__000001_NewDiagram_Event.NbInstances = 0
	__GongStructShape__000001_NewDiagram_Event.Width = 240.000000
	__GongStructShape__000001_NewDiagram_Event.Heigth = 93.000000
	__GongStructShape__000001_NewDiagram_Event.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_NewDiagram_GongsimCommand.Name = `NewDiagram-GongsimCommand`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimCommand]
	__GongStructShape__000002_NewDiagram_GongsimCommand.Identifier = `ref_models.GongsimCommand`
	__GongStructShape__000002_NewDiagram_GongsimCommand.ShowNbInstances = false
	__GongStructShape__000002_NewDiagram_GongsimCommand.NbInstances = 0
	__GongStructShape__000002_NewDiagram_GongsimCommand.Width = 240.000000
	__GongStructShape__000002_NewDiagram_GongsimCommand.Heigth = 138.000000
	__GongStructShape__000002_NewDiagram_GongsimCommand.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000003_NewDiagram_GongsimStatus.Name = `NewDiagram-GongsimStatus`
	
	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.GongsimStatus]
	__GongStructShape__000003_NewDiagram_GongsimStatus.Identifier = `ref_models.GongsimStatus`
	__GongStructShape__000003_NewDiagram_GongsimStatus.ShowNbInstances = false
	__GongStructShape__000003_NewDiagram_GongsimStatus.NbInstances = 0
	__GongStructShape__000003_NewDiagram_GongsimStatus.Width = 240.000000
	__GongStructShape__000003_NewDiagram_GongsimStatus.Heigth = 138.000000
	__GongStructShape__000003_NewDiagram_GongsimStatus.IsSelected = false

	// Position values setup
	__Position__000000_Pos_NewDiagram_ControlMode.X = 40.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Y = 590.000000
	__Position__000000_Pos_NewDiagram_ControlMode.Name = `Pos-NewDiagram-ControlMode`

	// Position values setup
	__Position__000001_Pos_NewDiagram_Engine.X = 74.000000
	__Position__000001_Pos_NewDiagram_Engine.Y = 89.000000
	__Position__000001_Pos_NewDiagram_Engine.Name = `Pos-NewDiagram-Engine`

	// Position values setup
	__Position__000002_Pos_NewDiagram_EngineRunMode.X = 28.000000
	__Position__000002_Pos_NewDiagram_EngineRunMode.Y = 43.000000
	__Position__000002_Pos_NewDiagram_EngineRunMode.Name = `Pos-NewDiagram-EngineRunMode`

	// Position values setup
	__Position__000003_Pos_NewDiagram_EngineState.X = 380.000000
	__Position__000003_Pos_NewDiagram_EngineState.Y = 80.000000
	__Position__000003_Pos_NewDiagram_EngineState.Name = `Pos-NewDiagram-EngineState`

	// Position values setup
	__Position__000004_Pos_NewDiagram_Event.X = 550.000000
	__Position__000004_Pos_NewDiagram_Event.Y = 320.000000
	__Position__000004_Pos_NewDiagram_Event.Name = `Pos-NewDiagram-Event`

	// Position values setup
	__Position__000005_Pos_NewDiagram_GongsimCommand.X = 70.000000
	__Position__000005_Pos_NewDiagram_GongsimCommand.Y = 330.000000
	__Position__000005_Pos_NewDiagram_GongsimCommand.Name = `Pos-NewDiagram-GongsimCommand`

	// Position values setup
	__Position__000006_Pos_NewDiagram_GongsimStatus.X = 540.000000
	__Position__000006_Pos_NewDiagram_GongsimStatus.Y = 470.000000
	__Position__000006_Pos_NewDiagram_GongsimStatus.Name = `Pos-NewDiagram-GongsimStatus`

	// Setup of pointers
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000000_NewDiagram_Engine)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000001_NewDiagram_Event)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000002_NewDiagram_GongsimCommand)
	__Classdiagram__000000_NewDiagram.GongStructShapes = append(__Classdiagram__000000_NewDiagram.GongStructShapes, __GongStructShape__000003_NewDiagram_GongsimStatus)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000002_NewDiagram_EngineState)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000000_NewDiagram_ControlMode)
	__Classdiagram__000000_NewDiagram.GongEnumShapes = append(__Classdiagram__000000_NewDiagram.GongEnumShapes, __GongEnumShape__000001_NewDiagram_EngineRunMode)
	__GongEnumShape__000000_NewDiagram_ControlMode.Position = __Position__000000_Pos_NewDiagram_ControlMode
	__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys, __GongEnumValueEntry__000000_AUTONOMOUS)
	__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys = append(__GongEnumShape__000000_NewDiagram_ControlMode.GongEnumValueEntrys, __GongEnumValueEntry__000001_CLIENT_CONTROL)
	__GongEnumShape__000001_NewDiagram_EngineRunMode.Position = __Position__000002_Pos_NewDiagram_EngineRunMode
	__GongEnumShape__000002_NewDiagram_EngineState.Position = __Position__000003_Pos_NewDiagram_EngineState
	__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000004_RUNNING)
	__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000003_PAUSED)
	__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys = append(__GongEnumShape__000002_NewDiagram_EngineState.GongEnumValueEntrys, __GongEnumValueEntry__000002_OVER)
	__GongStructShape__000000_NewDiagram_Engine.Position = __Position__000001_Pos_NewDiagram_Engine
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000013_Name)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000009_EndTime)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000006_CurrentTime)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000015_SecondsSinceStart)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000010_Fired)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000003_ControlMode)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000019_State)
	__GongStructShape__000000_NewDiagram_Engine.Fields = append(__GongStructShape__000000_NewDiagram_Engine.Fields, __Field__000016_Speed)
	__GongStructShape__000001_NewDiagram_Event.Position = __Position__000004_Pos_NewDiagram_Event
	__GongStructShape__000001_NewDiagram_Event.Fields = append(__GongStructShape__000001_NewDiagram_Event.Fields, __Field__000012_Name)
	__GongStructShape__000001_NewDiagram_Event.Fields = append(__GongStructShape__000001_NewDiagram_Event.Fields, __Field__000008_Duration)
	__GongStructShape__000002_NewDiagram_GongsimCommand.Position = __Position__000005_Pos_NewDiagram_GongsimCommand
	__GongStructShape__000002_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000002_NewDiagram_GongsimCommand.Fields, __Field__000011_Name)
	__GongStructShape__000002_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000002_NewDiagram_GongsimCommand.Fields, __Field__000000_Command)
	__GongStructShape__000002_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000002_NewDiagram_GongsimCommand.Fields, __Field__000001_CommandDate)
	__GongStructShape__000002_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000002_NewDiagram_GongsimCommand.Fields, __Field__000018_SpeedCommandType)
	__GongStructShape__000002_NewDiagram_GongsimCommand.Fields = append(__GongStructShape__000002_NewDiagram_GongsimCommand.Fields, __Field__000007_DateSpeedCommand)
	__GongStructShape__000003_NewDiagram_GongsimStatus.Position = __Position__000006_Pos_NewDiagram_GongsimStatus
	__GongStructShape__000003_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000003_NewDiagram_GongsimStatus.Fields, __Field__000014_Name)
	__GongStructShape__000003_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000003_NewDiagram_GongsimStatus.Fields, __Field__000004_CurrentCommand)
	__GongStructShape__000003_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000003_NewDiagram_GongsimStatus.Fields, __Field__000002_CompletionDate)
	__GongStructShape__000003_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000003_NewDiagram_GongsimStatus.Fields, __Field__000005_CurrentSpeedCommand)
	__GongStructShape__000003_NewDiagram_GongsimStatus.Fields = append(__GongStructShape__000003_NewDiagram_GongsimStatus.Fields, __Field__000017_SpeedCommandCompletionDate)
}


