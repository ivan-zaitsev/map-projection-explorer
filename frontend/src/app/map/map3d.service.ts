import {Injectable} from "@angular/core";
import {Map} from "ol";
// @ts-ignore
import OLCesium from "olcs";

@Injectable()
export class Map3dService {

  public createMap(map2d: Map, target: string): OLCesium {
    let map3d = new OLCesium({
      target: target,
      map: map2d,
      sceneOptions: {
        creditContainer: document.createElement("none")
      }
    });
    map3d.setEnabled(true);
    return map3d;
  }

}
