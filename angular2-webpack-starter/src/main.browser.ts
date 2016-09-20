/*
 * Angular bootstraping
 */
import { platformBrowserDynamic } from '@angular/platform-browser-dynamic';

/*
 * App Module
 * our top level module that holds all of our components
 */
import { AppModule } from './app';

document.addEventListener('DOMContentLoaded', () => {
  platformBrowserDynamic().bootstrapModule(AppModule).catch(err => console.error(err));
});

