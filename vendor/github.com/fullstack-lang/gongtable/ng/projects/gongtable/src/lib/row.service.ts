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

import { RowDB } from './row-db'
import { Row, CopyRowToRowDB } from './row'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { CellDB } from './cell-db'

@Injectable({
  providedIn: 'root'
})
export class RowService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  RowServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private rowsUrl: string

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
    this.rowsUrl = origin + '/api/github.com/fullstack-lang/gongtable/go/v1/rows';
  }

  /** GET rows from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB[]> {
    return this.getRows(GONG__StackPath, frontRepo)
  }
  getRows(GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<RowDB[]>(this.rowsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<RowDB[]>('getRows', []))
      );
  }

  /** GET row by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {
    return this.getRow(id, GONG__StackPath, frontRepo)
  }
  getRow(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.rowsUrl}/${id}`;
    return this.http.get<RowDB>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched row id=${id}`)),
      catchError(this.handleError<RowDB>(`getRow id=${id}`))
    );
  }

  // postFront copy row to a version with encoded pointers and post to the back
  postFront(row: Row, GONG__StackPath: string): Observable<RowDB> {
    let rowDB = new RowDB
    CopyRowToRowDB(row, rowDB)
    const id = typeof rowDB === 'number' ? rowDB : rowDB.ID
    const url = `${this.rowsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RowDB>(url, rowDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RowDB>('postRow'))
    );
  }
  
  /** POST: add a new row to the server */
  post(rowdb: RowDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {
    return this.postRow(rowdb, GONG__StackPath, frontRepo)
  }
  postRow(rowdb: RowDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<RowDB>(this.rowsUrl, rowdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted rowdb id=${rowdb.ID}`)
      }),
      catchError(this.handleError<RowDB>('postRow'))
    );
  }

  /** DELETE: delete the rowdb from the server */
  delete(rowdb: RowDB | number, GONG__StackPath: string): Observable<RowDB> {
    return this.deleteRow(rowdb, GONG__StackPath)
  }
  deleteRow(rowdb: RowDB | number, GONG__StackPath: string): Observable<RowDB> {
    const id = typeof rowdb === 'number' ? rowdb : rowdb.ID;
    const url = `${this.rowsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<RowDB>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted rowdb id=${id}`)),
      catchError(this.handleError<RowDB>('deleteRow'))
    );
  }

  // updateFront copy row to a version with encoded pointers and update to the back
  updateFront(row: Row, GONG__StackPath: string): Observable<RowDB> {
    let rowDB = new RowDB
    CopyRowToRowDB(row, rowDB)
    const id = typeof rowDB === 'number' ? rowDB : rowDB.ID
    const url = `${this.rowsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<RowDB>(url, rowDB, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<RowDB>('updateRow'))
    );
  }

  /** PUT: update the rowdb on the server */
  update(rowdb: RowDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {
    return this.updateRow(rowdb, GONG__StackPath, frontRepo)
  }
  updateRow(rowdb: RowDB, GONG__StackPath: string, frontRepo: FrontRepo): Observable<RowDB> {
    const id = typeof rowdb === 'number' ? rowdb : rowdb.ID;
    const url = `${this.rowsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<RowDB>(url, rowdb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated rowdb id=${rowdb.ID}`)
      }),
      catchError(this.handleError<RowDB>('updateRow'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in RowService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("RowService" + error); // log to console instead

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
