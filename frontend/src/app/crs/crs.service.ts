import {Injectable} from "@angular/core";
import {HttpClient, HttpParams} from "@angular/common/http";
import {CoordinateReferenceSystem, PagedList} from "./crs.model";
import {basePath} from "../app.config";
import {catchError, Observable, retry, throwError} from "rxjs";

@Injectable()
export class CrsService {

  constructor(private http: HttpClient) {
  }

  public findAll(pageSize: number, pageCursor?: number, search?: string):
    Observable<PagedList<CoordinateReferenceSystem>> {

    const url: string = basePath + "/api/v1/coordinate-reference-systems";

    let params = new HttpParams();
    if (search != null) {
      params = params.append("search", search);
    }
    if (pageCursor != null) {
      params = params.append("pageCursor", pageCursor);
    }
    params = params.append("pageSize", pageSize);

    return this.http.get<PagedList<CoordinateReferenceSystem>>(url, {params})
      .pipe(retry(1), catchError(error => throwError(() => error)));
  }

  public find(code: number): Observable<CoordinateReferenceSystem> {
    const url: string = basePath + `/api/v1/coordinate-reference-systems/${code}`;

    return this.http.get<CoordinateReferenceSystem>(url)
      .pipe(retry(1), catchError(error => throwError(() => error)));
  }

}
