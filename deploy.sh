#!/bin/bash

mkdir -p out
bash build.sh
cp prompt/* out/
cp example.env out/
cd out
mv example.env .env
nano .env
