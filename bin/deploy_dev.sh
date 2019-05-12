#!/usr/bin/env bash
# Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT license
# that can be found in the LICENSE file.

echo '[dev] Applying environment variables and configs...'
cp .env.dev.dist .env
cp Dockerfile.dev.dist Dockerfile
cp docker-compose.dev.yml.dist docker-compose.yml
source .env
echo

echo "[${APP_ENV}] Stopping docker-compose services..."
docker-compose down
echo

echo "[${APP_ENV}] Building docker-compose services..."
docker-compose build --force-rm
echo

echo "[${APP_ENV}] Starting docker-compose services..."
docker-compose up -d
echo

echo "[${APP_ENV}] Clearing old environment..."
docker rm $(docker ps -qa --no-trunc --filter "status=exited")
docker rmi $(docker images -f "dangling=true" -q)
