import {Injectable} from "@angular/core";
import {ProjectionLike, transform, transformExtent} from "ol/proj";
import {Map, View} from "ol";
import {Tile as TileLayer, Vector as VectorLayer} from "ol/layer";
import {OSM, Vector as VectorSource} from "ol/source";
import Polygon from "ol/geom/Polygon";
import {Extent} from "ol/extent";
import Feature from "ol/Feature";
import {Geometry} from "ol/geom";
import {Fill, Stroke, Style} from "ol/style";
import {CoordinateReferenceSystemExtent, CoordinateReferenceSystemExtentCenter} from "../crs/crs.model";

@Injectable()
export class Map2dService {

  public static readonly PROJECTION_CODE_PREFIX: string = "EPSG:"
  public static readonly ORIGINAL_PROJECTION_CODE: string = Map2dService.PROJECTION_CODE_PREFIX + "4326"

  public createMap(
    projectionCenter: number[], projection: ProjectionLike, projectionExtent: Extent, target?: string): Map {
    let map = new Map({
      target: target,
      layers: [
        new TileLayer({
          source: new OSM(),
        })
      ],
      view: new View({
        center: projectionCenter,
        projection: projection,
        showFullExtent: true
      }),
    });

    map.getView().fit(projectionExtent)
    return map;
  }

  public buildPolygon(extent: Extent): Polygon {
    return new Polygon([
      [
        [extent[0], extent[1]],
        [extent[0], extent[3]],
        [extent[2], extent[3]],
        [extent[2], extent[1]],
        [extent[0], extent[1]],
      ],
    ]);
  }

  public buildPolygonLayer(polygonSource: VectorSource): VectorLayer<Feature<Geometry>> {
    return new VectorLayer({
      source: polygonSource,
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
  }

  public transformExtent(extent: CoordinateReferenceSystemExtent, toCode: string): Extent {
    return transformExtent(
      [extent.westLongitude, extent.southLatitude, extent.eastLongitude, extent.northLatitude],
      Map2dService.ORIGINAL_PROJECTION_CODE,
      toCode,
    );
  }

  public transformExtentCenter(extentCenter: CoordinateReferenceSystemExtentCenter, toCode: string): Extent {
    return transform(
      [extentCenter.longitude, extentCenter.latitude],
      Map2dService.ORIGINAL_PROJECTION_CODE,
      toCode,
    );
  }

}
