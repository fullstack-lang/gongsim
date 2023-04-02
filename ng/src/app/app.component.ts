import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongsim from 'gongsim'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongsim Data/Model'
  sim = 'Engine Control'
  view = this.sim

  views: string[] = [this.sim, this.default];

  GONG__StackPath = "github.com/fullstack-lang/gongsim/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}
