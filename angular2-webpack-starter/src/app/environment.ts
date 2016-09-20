import { enableDebugTools, disableDebugTools } from '@angular/platform-browser';
import { enableProdMode, ApplicationRef } from '@angular/core';
// Environment Providers
let PROVIDERS: any[] = [
  // common env directives
];

if ('production' === ENV) {
  disableDebugTools();
  enableProdMode();
  PROVIDERS = [
    ...PROVIDERS
  ];
} else {
  // DEVELOPMENT
  PROVIDERS = [
    ...PROVIDERS
  ];
}

export const ENV_PROVIDERS = [
  ...PROVIDERS
];
