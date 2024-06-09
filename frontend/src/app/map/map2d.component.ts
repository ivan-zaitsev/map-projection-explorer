import {afterNextRender, Component, OnInit} from '@angular/core';
import proj4 from 'proj4';
import {register} from 'ol/proj/proj4';
import {get as getProjection, ProjectionLike} from 'ol/proj';
import {Map} from 'ol';
import {Vector as VectorSource} from 'ol/source';
import Feature from 'ol/Feature';
import {Map2dService} from "./map2d.service";
import {CrsService} from "../crs/crs.service";
import {ActivatedRoute} from "@angular/router";
import {Observable} from "rxjs";
import {CoordinateReferenceSystem} from "../crs/crs.model";

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [],
  providers: [
    CrsService,
    Map2dService
  ],
  templateUrl: './map.component.html',
  styleUrl: './map.component.scss'
})
export class Map2dComponent implements OnInit {

  code?: number

  map2d!: Map
  map2dPolygonSource!: VectorSource

  public constructor(
    private route: ActivatedRoute,
    private crsService: CrsService,
    private map2dService: Map2dService) {

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
    const result: Observable<CoordinateReferenceSystem> = this.crsService.find(code);

    result.subscribe({
      next: (data: CoordinateReferenceSystem) => {
        const code: string = Map2dService.PROJECTION_CODE_PREFIX + data.code;
        this.registerProjection(code, data.definition.value);

        const projection = getProjection(code) as ProjectionLike
        const projectionExtent = this.map2dService.transformExtent(data.extent, code);
        const projectionExtentCenter = this.map2dService.transformExtentCenter(data.extentCenter, code);

        this.map2d = this.map2dService.createMap(projectionExtentCenter, projection, projectionExtent, 'map');

        this.map2dPolygonSource = new VectorSource();
        this.map2d.addLayer(this.map2dService.buildPolygonLayer(this.map2dPolygonSource));
        this.map2dPolygonSource.addFeature(new Feature(this.map2dService.buildPolygon(projectionExtent)));
      },
      error: (error: any) => console.log(error)
    });
  }

  private registerProjection(code: string, proj4text: string): void {
    proj4.defs(code, proj4text,);
    register(proj4);
  }

}
