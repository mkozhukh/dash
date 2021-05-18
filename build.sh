#!/bin/bash
set -e

cd app && yarn && yarn build && cd ..
mkdir -p app/public/dist
mkdir -p app/public/assets
cp -rf app/public/dist/*.{js,css} backend/public/dist/
cp -rf app/public/assets/* backend/public/assets/
cp -rf app/public/index.production.html backend/public/index.html
cd backend && pkger && go build
