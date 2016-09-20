import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { RouterModule } from '@angular/router';
import { UniversalModule } from 'angular2-universal';

import { App } from './app/app';
import { HomeComponent } from './app/home/home.component';
import { routing } from './app/app.routing';

@NgModule({
  bootstrap: [App],
  declarations: [
    App, HomeComponent
  ],
  imports: [
    UniversalModule, // NodeModule, NodeHttpModule, and NodeJsonpModule are included
    FormsModule,
    routing
  ]
})
export class MainModule {

}