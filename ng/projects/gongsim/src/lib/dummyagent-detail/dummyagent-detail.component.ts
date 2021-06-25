// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { DummyAgentDB } from '../dummyagent-db'
import { DummyAgentService } from '../dummyagent.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-dummyagent-detail',
	templateUrl: './dummyagent-detail.component.html',
	styleUrls: ['./dummyagent-detail.component.css'],
})
export class DummyAgentDetailComponent implements OnInit {

	// insertion point for declarations

	// the DummyAgentDB of interest
	dummyagent: DummyAgentDB;

	// front repo
	frontRepo: FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	constructor(
		private dummyagentService: DummyAgentService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getDummyAgent()

		// observable for changes in structs
		this.dummyagentService.DummyAgentServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getDummyAgent()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getDummyAgent(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.dummyagent = frontRepo.DummyAgents.get(id)
				} else {
					this.dummyagent = new (DummyAgentDB)
				}

				// insertion point for recovery of form controls value for bool fields
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		if (this.dummyagent.EngineID == undefined) {
			this.dummyagent.EngineID = new NullInt64
		}
		if (this.dummyagent.Engine != undefined) {
			this.dummyagent.EngineID.Int64 = this.dummyagent.Engine.ID
			this.dummyagent.EngineID.Valid = true
		} else {
			this.dummyagent.EngineID.Int64 = 0
			this.dummyagent.EngineID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.dummyagentService.updateDummyAgent(this.dummyagent)
				.subscribe(dummyagent => {
					this.dummyagentService.DummyAgentServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.dummyagentService.postDummyAgent(this.dummyagent).subscribe(dummyagent => {

				this.dummyagentService.DummyAgentServiceChanged.next("post")

				this.dummyagent = {} // reset fields
			});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		dialogConfig.data = {
			ID: this.dummyagent.ID,
			ReversePointer: reverseField,
			OrderingMode: false,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'sTableComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.dummyagent.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event) {
		if (this.dummyagent.Name == undefined) {
			this.dummyagent.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}
