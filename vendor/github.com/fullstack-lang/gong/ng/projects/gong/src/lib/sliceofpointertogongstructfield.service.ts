// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule, HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { SliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield-db'
import { SliceOfPointerToGongStructField, CopySliceOfPointerToGongStructFieldToSliceOfPointerToGongStructFieldDB } from './sliceofpointertogongstructfield'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { GongStructDB } from './gongstruct-db'

@Injectable({
  providedIn: 'root'
})
export class SliceOfPointerToGongStructFieldService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  SliceOfPointerToGongStructFieldServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private sliceofpointertogongstructfieldsUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.sliceofpointertogongstructfieldsUrl = origin + '/api/github.com/fullstack-lang/gong/go/v1/sliceofpointertogongstructfields';
  }

  /** GET sliceofpointertogongstructfields from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB[]> {
    return this.getSliceOfPointerToGongStructFields(GONG__StackPath, frontRepo)
  }
  getSliceOfPointerToGongStructFields(GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<SliceOfPointerToGongStructFieldDB[]>(this.sliceofpointertogongstructfieldsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<SliceOfPointerToGongStructFieldDB[]>('getSliceOfPointerToGongStructFields', []))
      );
  }

  /** GET sliceofpointertogongstructfield by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {
    return this.getSliceOfPointerToGongStructField(id, GONG__StackPath, frontRepo)
  }
  getSliceOfPointerToGongStructField(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;
    return this.http.get<SliceOfPointerToGongStructFieldDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched sliceofpointertogongstructfield id=${id}`)),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>(`getSliceOfPointerToGongStructField id=${id}`))
    );
  }

  // postFront copy sliceofpointertogongstructfield to a version with encoded pointers and post to the back
  postFront(sliceofpointertogongstructfield: SliceOfPointerToGongStructField, GONG__StackPath: string): Observable<SliceOfPointerToGongStructFieldDB> {
    let sliceofpointertogongstructfieldDB = new SliceOfPointerToGongStructFieldDB
    CopySliceOfPointerToGongStructFieldToSliceOfPointerToGongStructFieldDB(sliceofpointertogongstructfield, sliceofpointertogongstructfieldDB)
    const id = typeof sliceofpointertogongstructfieldDB === 'number' ? sliceofpointertogongstructfieldDB : sliceofpointertogongstructfieldDB.ID
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SliceOfPointerToGongStructFieldDB>(url, sliceofpointertogongstructfieldDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('postSliceOfPointerToGongStructField'))
    );
  }
  
  /** POST: add a new sliceofpointertogongstructfield to the server */
  post(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {
    return this.postSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb, GONG__StackPath, frontRepo)
  }
  postSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<SliceOfPointerToGongStructFieldDB>(this.sliceofpointertogongstructfieldsUrl, sliceofpointertogongstructfielddb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted sliceofpointertogongstructfielddb id=${sliceofpointertogongstructfielddb.ID}`)
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('postSliceOfPointerToGongStructField'))
    );
  }

  /** DELETE: delete the sliceofpointertogongstructfielddb from the server */
  delete(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB | number, GONG__StackPath: string): Observable<SliceOfPointerToGongStructFieldDB> {
    return this.deleteSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb, GONG__StackPath)
  }
  deleteSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB | number, GONG__StackPath: string): Observable<SliceOfPointerToGongStructFieldDB> {
    const id = typeof sliceofpointertogongstructfielddb === 'number' ? sliceofpointertogongstructfielddb : sliceofpointertogongstructfielddb.ID;
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<SliceOfPointerToGongStructFieldDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted sliceofpointertogongstructfielddb id=${id}`)),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('deleteSliceOfPointerToGongStructField'))
    );
  }

  // updateFront copy sliceofpointertogongstructfield to a version with encoded pointers and update to the back
  updateFront(sliceofpointertogongstructfield: SliceOfPointerToGongStructField, GONG__StackPath: string): Observable<SliceOfPointerToGongStructFieldDB> {
    let sliceofpointertogongstructfieldDB = new SliceOfPointerToGongStructFieldDB
    CopySliceOfPointerToGongStructFieldToSliceOfPointerToGongStructFieldDB(sliceofpointertogongstructfield, sliceofpointertogongstructfieldDB)
    const id = typeof sliceofpointertogongstructfieldDB === 'number' ? sliceofpointertogongstructfieldDB : sliceofpointertogongstructfieldDB.ID
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<SliceOfPointerToGongStructFieldDB>(url, sliceofpointertogongstructfieldDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('updateSliceOfPointerToGongStructField'))
    );
  }

  /** PUT: update the sliceofpointertogongstructfielddb on the server */
  update(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {
    return this.updateSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb, GONG__StackPath, frontRepo)
  }
  updateSliceOfPointerToGongStructField(sliceofpointertogongstructfielddb: SliceOfPointerToGongStructFieldDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<SliceOfPointerToGongStructFieldDB> {
    const id = typeof sliceofpointertogongstructfielddb === 'number' ? sliceofpointertogongstructfielddb : sliceofpointertogongstructfielddb.ID;
    const url = `${this.sliceofpointertogongstructfieldsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<SliceOfPointerToGongStructFieldDB>(url, sliceofpointertogongstructfielddb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated sliceofpointertogongstructfielddb id=${sliceofpointertogongstructfielddb.ID}`)
      }),
      catchError(this.handleError<SliceOfPointerToGongStructFieldDB>('updateSliceOfPointerToGongStructField'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in SliceOfPointerToGongStructFieldService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("SliceOfPointerToGongStructFieldService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}
