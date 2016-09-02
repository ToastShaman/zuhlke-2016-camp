#!/bin/sh

rm -rf ./dist
mkdir ./dist

for dir in ./src/*
do
    dir=${dir%*/}
    npm run tsc -- --outDir ./dist/${dir##*/} src/${dir##*/}/main.ts
done
