// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { GongsimCommandDB } from '../gongsimcommand-db'
import { GongsimCommandService } from '../gongsimcommand.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { GongsimCommandTypeSelect, GongsimCommandTypeList } from '../GongsimCommandType'
import { SpeedCommandTypeSelect, SpeedCommandTypeList } from '../SpeedCommandType'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../front-repo.service'

@Component({
	selector: 'app-gongsimcommand-detail',
	templateUrl: './gongsimcommand-detail.component.html',
	styleUrls: ['./gongsimcommand-detail.component.css'],
})
export class GongsimCommandDetailComponent implements OnInit {

	// insertion point for declarations
	GongsimCommandTypeList: GongsimCommandTypeSelect[]
	SpeedCommandTypeList: SpeedCommandTypeSelect[]

	// the GongsimCommandDB of interest
	gongsimcommand: GongsimCommandDB;

	// front repo
	frontRepo: FrontRepo

	constructor(
		private gongsimcommandService: GongsimCommandService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {
		this.getGongsimCommand()

		// observable for changes in structs
		this.gongsimcommandService.GongsimCommandServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getGongsimCommand()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.GongsimCommandTypeList = GongsimCommandTypeList
		this.SpeedCommandTypeList = SpeedCommandTypeList
	}

	getGongsimCommand(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		const association = this.route.snapshot.paramMap.get('association');

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo
				if (id != 0 && association == undefined) {
					this.gongsimcommand = frontRepo.GongsimCommands.get(id)
				} else {
					this.gongsimcommand = new (GongsimCommandDB)
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
		
		// save from the front pointer space to the non pointer space for serialization
		if (association == undefined) {
			// insertion point for translation/nullation of each pointers
		}

		if (id != 0 && association == undefined) {

			this.gongsimcommandService.updateGongsimCommand(this.gongsimcommand)
				.subscribe(gongsimcommand => {
					this.gongsimcommandService.GongsimCommandServiceChanged.next("update")
				});
		} else {
			switch (association) {
				// insertion point for saving value of ONE_MANY association reverse pointer
			}
			this.gongsimcommandService.postGongsimCommand(this.gongsimcommand).subscribe(gongsimcommand => {

				this.gongsimcommandService.GongsimCommandServiceChanged.next("post")

				this.gongsimcommand = {} // reset fields
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
			ID: this.gongsimcommand.ID,
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
			ID: this.gongsimcommand.ID,
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
}
