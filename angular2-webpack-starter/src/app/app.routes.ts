import { Routes, RouterModule } from '@angular/router';
import { NoContent } from './no-content';
import { Hello } from './hello/hello.component';

export const ROUTES: Routes = [
  { path: '',      component: NoContent },
  { path: 'hello', component: Hello },
  { path: '**',    component: NoContent }
];
