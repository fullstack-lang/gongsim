// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { EventDB } from '../event-db'
import { EventService } from '../event.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-event-sorting',
  templateUrl: './event-sorting.component.html',
  styleUrls: ['./event-sorting.component.css']
})
export class EventSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Event instances that are in the association
  associatedEvents = new Array<EventDB>();

  constructor(
    private eventService: EventService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of event instances
    public dialogRef: MatDialogRef<EventSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getEvents()
  }

  getEvents(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let event of this.frontRepo.Events_array) {
          let ID = this.dialogData.ID
          let revPointerID = event[this.dialogData.ReversePointer as keyof EventDB] as unknown as NullInt64
          let revPointerID_Index = event[this.dialogData.ReversePointer + "_Index" as keyof EventDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedEvents.push(event)
          }
        }

        // sort associated event according to order
        this.associatedEvents.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedEvents, event.previousIndex, event.currentIndex);

    // set the order of Event instances
    let index = 0

    for (let event of this.associatedEvents) {
      let revPointerID_Index = event[this.dialogData.ReversePointer + "_Index" as keyof EventDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedEvents.forEach(
      event => {
        this.eventService.updateEvent(event)
          .subscribe(event => {
            this.eventService.EventServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}
