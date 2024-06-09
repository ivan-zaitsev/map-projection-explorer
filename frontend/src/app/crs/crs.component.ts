import {Component, OnInit} from '@angular/core';
import {Observable} from "rxjs";
import {CoordinateReferenceSystem, PagedList} from "./crs.model";
import {CrsService} from "./crs.service";
import {NgForOf} from "@angular/common";
import {RouterLink, RouterLinkActive} from "@angular/router";

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [NgForOf, RouterLink, RouterLinkActive],
  providers: [CrsService],
  templateUrl: './crs.component.html',
  styleUrl: './crs.component.scss'
})
export class CrsComponent implements OnInit {

  search?: string
  pageSize: number = 50;
  nextPageCursor?: number;
  coordinateReferenceSystems!: CoordinateReferenceSystem[];

  constructor(private crsService: CrsService) {
  }

  public ngOnInit(): void {
    this.coordinateReferenceSystems = [];
    this.render(this.pageSize, undefined, undefined);
  }

  public onSearch(search: string): void {
    this.search = search;
    this.nextPageCursor = undefined
    this.coordinateReferenceSystems = [];
    this.render(this.pageSize, undefined, this.search);
  }

  public onLoadMore(): void {
    this.render(this.pageSize, this.nextPageCursor, this.search);
  }

  public render(pageSize: number, pageCursor?: number, search?: string): void {
    const result: Observable<PagedList<CoordinateReferenceSystem>> =
      this.crsService.findAll(pageSize, pageCursor, search);

    result.subscribe({
      next: (data: PagedList<CoordinateReferenceSystem>) => {
        this.coordinateReferenceSystems.push(...data.content);
        this.nextPageCursor = data.nextPageCursor;
      },
      error: (error: any) => console.log(error)
    });
  }

}
