import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { GongsimModule } from 'gongsim'
import { GongsimcontrolModule } from 'gongsimcontrol'

import { AngularSplitModule, SplitComponent } from 'angular-split';

// mandatory
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    AngularSplitModule,

    HttpClientModule,
    GongsimModule,
    GongsimcontrolModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
