// insertion point sub template for components imports 
import { DummyAgentsTableComponent } from './dummyagents-table/dummyagents-table.component'
import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EventsTableComponent } from './events-table/events-table.component'
import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
import { UpdateStatesTableComponent } from './updatestates-table/updatestates-table.component'

// insertion point sub template for map of components per struct 
export const MapOfDummyAgentsComponents: Map<string, any> = new Map([["DummyAgentsTableComponent", DummyAgentsTableComponent],])
export const MapOfEnginesComponents: Map<string, any> = new Map([["EnginesTableComponent", EnginesTableComponent],])
export const MapOfEventsComponents: Map<string, any> = new Map([["EventsTableComponent", EventsTableComponent],])
export const MapOfGongsimCommandsComponents: Map<string, any> = new Map([["GongsimCommandsTableComponent", GongsimCommandsTableComponent],])
export const MapOfGongsimStatussComponents: Map<string, any> = new Map([["GongsimStatussTableComponent", GongsimStatussTableComponent],])
export const MapOfUpdateStatesComponents: Map<string, any> = new Map([["UpdateStatesTableComponent", UpdateStatesTableComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["DummyAgent", MapOfDummyAgentsComponents],
      ["Engine", MapOfEnginesComponents],
      ["Event", MapOfEventsComponents],
      ["GongsimCommand", MapOfGongsimCommandsComponents],
      ["GongsimStatus", MapOfGongsimStatussComponents],
      ["UpdateState", MapOfUpdateStatesComponents],
    ]
  )