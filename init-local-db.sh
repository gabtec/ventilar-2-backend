#!/bin/bash

PG_CT_VERSION="postgres:15.1-alpine"

echo "[ INFO:] Deploying local databases from .env, using docker..."
echo "[ INFO:] Postgresql container version: $PG_CT_VERSION"

# if .env not found exit
if [ ! -e ".env" ]; then
  echo "[ERROR:] File \".env\" not found. Create one from .env.example. Exiting."
  exit 1
fi

DB_NAME=$(cat ./.env | grep DB_NAME | cut -d "=" -f 2)
DB_USER=$(cat ./.env | grep DB_USER | cut -d "=" -f 2)
DB_SECRET=$(cat ./.env | grep DB_SECRET | cut -d "=" -f 2)
DB_PORT=$(cat ./.env | grep DB_PORT | cut -d "=" -f 2)

CT_NAME=postgres-ct

echo "[ INFO:] Deploying a database container..."

# check PORT is free
nc -z localhost $DB_PORT

if [ $? == 0 ]; then
  # PORT in use
  echo "PORT $DB_PORT is in use. Must select another one."
  exit 4
fi 

# Deploy database container
docker run --name $CT_NAME -e POSTGRES_USER=$DB_USER -e POSTGRES_PASSWORD=$DB_SECRET -e POSTGRES_DB=$DB_NAME  -p $DB_PORT:5432 -d $PG_CT_VERSION

sleep 1

# # IP=$(docker container inspect $CT_NAME -f '{{ .NetworkSettings.Networks.bridge.IPAddress }}')
pingpong() {
  RES=$(docker exec -it postgres-ct psql -U admin --no-password atlas_db -c "\l");
  return $?;
}

while ! pingpong ;do
  echo -n "..."
  sleep 1
done 

echo " "
echo "[ INFO:] Setting up credentials..."
docker exec -it $CT_NAME bash -c "echo $DB_SECRET > /root/.pgpass && chmod 0600 /root/.pgpass"

echo "[ INFO:] Creating auxiliary databases..."
docker exec -it $CT_NAME psql -U $DB_USER --no-password $DB_NAME -c "CREATE DATABASE dev;" > /dev/null 2>&1  # for atlas.io
docker exec -it $CT_NAME psql -U $DB_USER --no-password $DB_NAME -c "CREATE DATABASE ${DB_NAME}_test;" > /dev/null 2>&1 # for tests
echo "[ INFO:] Databases \"dev\" and \"${DB_NAME}_test\" created."

echo "[   OK:] Done."

# Test command:
# docker exec -it postgres-ct psql -U admin --no-password atlas_db -c "\l"  # for atlas.io