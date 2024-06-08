import {afterNextRender, Component} from '@angular/core';
import proj4 from 'proj4';
import {register} from 'ol/proj/proj4';
import {get as getProjection, ProjectionLike, transformExtent} from 'ol/proj';
import {Map, View} from 'ol';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer';
import {OSM, Vector as VectorSource} from 'ol/source';
import Feature from 'ol/Feature';
import Polygon from 'ol/geom/Polygon';
import {Fill, Stroke, Style} from 'ol/style';

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [],
  templateUrl: './map.component.html',
  styleUrl: './map.component.scss'
})
export class Map2dComponent {

  map2d!: Map

  public constructor() {
    afterNextRender(() => {
      proj4.defs(
        'EPSG:3035',
        '+proj=laea +lat_0=52 +lon_0=10 +x_0=4321000 +y_0=3210000 +ellps=GRS80 +units=m +no_defs',
      );
      register(proj4);

      const projection = getProjection('EPSG:3035');
      const projectionCenter = [4321000, 3210000];

      const projectionBorders = transformExtent(
        [-35.58, 24.6, 44.83, 84.73],
        'EPSG:4326',
        'EPSG:3035',
      );

      const polygon = new Polygon([
        [
          [projectionBorders[0], projectionBorders[1]],
          [projectionBorders[0], projectionBorders[3]],
          [projectionBorders[2], projectionBorders[3]],
          [projectionBorders[2], projectionBorders[1]],
          [projectionBorders[0], projectionBorders[1]],
        ],
      ]);

      const projectionBordersFeature = new Feature(polygon);
      const vectorSource = new VectorSource({
        features: [projectionBordersFeature],
      });

      const vectorLayer = new VectorLayer({
        source: vectorSource,
        style: new Style({
          stroke: new Stroke({
            color: 'red',
            width: 2,
          }),
          fill: new Fill({
            color: 'rgba(255, 0, 0, 0.1)',
          }),
        }),
      });

      this.map2d = this.buildMap2d(projectionCenter, projection as ProjectionLike);
      this.map2d.addLayer(vectorLayer);
    });
  }

  private buildMap2d(projectionCenter: number[], projection: ProjectionLike) {
    return new Map({
      target: 'map',
      layers: [
        new TileLayer({
          source: new OSM(),
        })
      ],
      view: new View({
        zoom: 2,
        center: projectionCenter,
        projection: projection,
        showFullExtent: true
      }),
    });
  }

}
