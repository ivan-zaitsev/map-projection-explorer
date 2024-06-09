import {afterNextRender, Component, OnInit} from '@angular/core';
import {get as getProjection, ProjectionLike} from 'ol/proj';
import {Map} from 'ol';
// @ts-ignore
import OLCesium from 'olcs';
import {Map3dService} from "./map3d.service";
import {Map2dService} from "./map2d.service";
import {Vector as VectorSource} from "ol/source";
import Feature from "ol/Feature";
import {ActivatedRoute} from "@angular/router";
import {CrsService} from "../crs/crs.service";
import {Observable} from "rxjs";
import {CoordinateReferenceSystem} from "../crs/crs.model";

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [],
  providers: [
    CrsService,
    Map2dService,
    Map3dService,
  ],
  templateUrl: './map.component.html',
  styleUrl: './map.component.scss'
})
export class Map3dComponent implements OnInit {

  code?: number

  map2d!: Map
  map2dPolygonSource!: VectorSource
  map3d!: OLCesium

  public constructor(
    private route: ActivatedRoute,
    private crsService: CrsService,
    private map2dService: Map2dService,
    private map3dService: Map3dService) {

    afterNextRender(() => {
      if (this.code !== undefined) {
        this.render(this.code)
      }
    });
  }

  async ngOnInit(): Promise<void> {
    await new Promise<void>(() => {
      this.route.queryParams.subscribe(params => {
        this.code = params["code"];
      });
    });
  }

  public render(code: number): void {
    console.log(code);

    const result: Observable<CoordinateReferenceSystem> = this.crsService.find(code);

    result.subscribe({
      next: (data: CoordinateReferenceSystem) => {
        const code = Map2dService.ORIGINAL_PROJECTION_CODE;
        const projection = getProjection(code) as ProjectionLike;
        const projectionExtent = this.map2dService.transformExtent(data.extent, code);
        const projectionExtentCenter = this.map2dService.transformExtentCenter(data.extentCenter, code);

        this.map2d = this.map2dService.createMap(projectionExtentCenter, projection, projectionExtent);
        this.map3d = this.map3dService.createMap(this.map2d, 'map');

        this.map2dPolygonSource = new VectorSource();
        this.map2d.addLayer(this.map2dService.buildPolygonLayer(this.map2dPolygonSource));
        this.map2dPolygonSource.addFeature(new Feature(this.map2dService.buildPolygon(projectionExtent)));
      },
      error: (error: any) => console.log(error)
    });
  }

}
