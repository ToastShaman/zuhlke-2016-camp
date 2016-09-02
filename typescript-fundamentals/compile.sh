#!/bin/sh

rm -rf ./dist
mkdir ./dist
npm run tsc -- --outDir ./dist/example-01 src/example-01/main.ts
