import {afterNextRender, Component} from '@angular/core';
import {get as getProjection, ProjectionLike, transformExtent} from 'ol/proj';
import {Map, View} from 'ol';
import {Tile as TileLayer, Vector as VectorLayer} from 'ol/layer';
import {OSM, Vector as VectorSource} from 'ol/source';
import Feature from 'ol/Feature';
import Polygon from 'ol/geom/Polygon';
import {Fill, Stroke, Style} from 'ol/style';

// @ts-ignore
import OLCesium from 'olcs';

@Component({
  selector: 'app-map',
  standalone: true,
  imports: [],
  templateUrl: './map.component.html',
  styleUrl: './map.component.scss'
})
export class Map3dComponent {

  map2d!: Map
  map3d!: OLCesium

  public constructor() {
    afterNextRender(() => {
      const projection = getProjection('EPSG:4326');
      const projectionCenter = [4321000, 3210000];

      const projectionBorders = transformExtent(
        [-35.58, 24.6, 44.83, 84.73],
        'EPSG:4326',
        'EPSG:4326',
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
      this.map3d = this.buildMap3d(this.map2d);

      this.map2d.addLayer(vectorLayer);
    });
  }

  private buildMap2d(projectionCenter: number[], projection: ProjectionLike) {
    return new Map({
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

  private buildMap3d(map2d: Map): OLCesium {
    var map3d = new OLCesium({
      target: 'map',
      map: map2d,
      sceneOptions: {
        creditContainer: document.createElement("none")
      }
    });
    map3d.setEnabled(true);
    return map3d;
  }

}
