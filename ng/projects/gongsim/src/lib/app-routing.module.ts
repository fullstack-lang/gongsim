import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { DummyAgentsTableComponent } from './dummyagents-table/dummyagents-table.component'
import { DummyAgentDetailComponent } from './dummyagent-detail/dummyagent-detail.component'

import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'

import { EventsTableComponent } from './events-table/events-table.component'
import { EventDetailComponent } from './event-detail/event-detail.component'

import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
import { GongsimCommandDetailComponent } from './gongsimcommand-detail/gongsimcommand-detail.component'

import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
import { GongsimStatusDetailComponent } from './gongsimstatus-detail/gongsimstatus-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagents', component: DummyAgentsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-adder', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-adder/:id/:originStruct/:originStructFieldName', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-detail/:id', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-engines', component: EnginesTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-adder', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-adder/:id/:originStruct/:originStructFieldName', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-detail/:id', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-events', component: EventsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-adder', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-adder/:id/:originStruct/:originStructFieldName', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-detail/:id', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommands', component: GongsimCommandsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-adder', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-adder/:id/:originStruct/:originStructFieldName', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-detail/:id', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatuss', component: GongsimStatussTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-adder', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-adder/:id/:originStruct/:originStructFieldName', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-detail/:id', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
