export interface PagedList<T> {

  content: T[];
  nextPageCursor: number;

}

export interface CoordinateReferenceSystem {
  name: string;
  code: number;
  definition: CoordinateReferenceSystemDefinition;
  extent: CoordinateReferenceSystemExtent;
  extentCenter: CoordinateReferenceSystemExtentCenter;
}

export interface CoordinateReferenceSystemDefinition {
  type: string;
  value: string;
}

export interface CoordinateReferenceSystemExtent {
  westLongitude: number;
  southLatitude: number;
  eastLongitude: number;
  northLatitude: number;
}

export interface CoordinateReferenceSystemExtentCenter {
  longitude: number;
  latitude: number;
}
