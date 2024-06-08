import {ApplicationConfig, provideZoneChangeDetection} from '@angular/core';
import {provideRouter, Routes} from '@angular/router';

import {provideClientHydration} from '@angular/platform-browser';
import {Map2dComponent} from './map/map2d.component';
import {Map3dComponent} from './map/map3d.component';

const routes: Routes = [
  {path: 'map2d', component: Map2dComponent},
  {path: 'map3d', component: Map3dComponent}
];

export const appConfig: ApplicationConfig = {
  providers: [
    provideZoneChangeDetection({eventCoalescing: true}),
    provideRouter(routes),
    provideClientHydration()
  ]
};
