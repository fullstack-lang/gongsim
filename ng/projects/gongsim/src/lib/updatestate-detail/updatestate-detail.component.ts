// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { UpdateStateDB } from '../updatestate-db'
import { UpdateStateService } from '../updatestate.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-updatestate-detail',
	templateUrl: './updatestate-detail.component.html',
	styleUrls: ['./updatestate-detail.component.css'],
})
export class UpdateStateDetailComponent implements OnInit {

	// insertion point for declarations
	Duration_Hours: number
	Duration_Minutes: number
	Duration_Seconds: number
	Period_Hours: number
	Period_Minutes: number
	Period_Seconds: number

	// the UpdateStateDB of interest
	updatestate: UpdateStateDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private updatestateService: UpdateStateService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getUpdateState()

		// observable for changes in structs
		this.updatestateService.UpdateStateServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getUpdateState()
				}
			}
		)

		// insertion point for initialisation of enums list
	}

	getUpdateState(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.updatestate = frontRepo.UpdateStates.get(id)
				} else {
					this.updatestate = new (UpdateStateDB)
				}

				// insertion point for recovery of form controls value for bool fields
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.updatestate.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.updatestate.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.updatestate.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
				// computation of Hours, Minutes, Seconds for Period
				this.Period_Hours = Math.floor(this.updatestate.Period / (3600 * 1000 * 1000 * 1000))
				this.Period_Minutes = Math.floor(this.updatestate.Period % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Period_Seconds = this.updatestate.Period % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization
		
		// insertion point for translation/nullation of each field
		this.updatestate.Duration =
			this.Duration_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Duration_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Duration_Seconds * (1000 * 1000 * 1000)
		this.updatestate.Period =
			this.Period_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Period_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Period_Seconds * (1000 * 1000 * 1000)
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.updatestateService.updateUpdateState(this.updatestate)
				.subscribe(updatestate => {
					this.updatestateService.UpdateStateServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.updatestateService.postUpdateState(this.updatestate).subscribe(updatestate => {

				this.updatestateService.UpdateStateServiceChanged.next("post")

				this.updatestate = {} // reset fields
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
			ID: this.updatestate.ID,
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
			ID: this.updatestate.ID,
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
		if (this.updatestate.Name == undefined) {
			this.updatestate.Name = event.value.Name		
		}
	}
}
