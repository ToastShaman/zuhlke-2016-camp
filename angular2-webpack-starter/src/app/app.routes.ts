import { Routes, RouterModule } from '@angular/router';
import { NoContent } from './no-content';

export const ROUTES: Routes = [
  { path: '',      component: NoContent },
  { path: '**',    component: NoContent }
];
