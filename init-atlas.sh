#!/bin/bash

echo "[ INFO:] Creating database folder structure..."

mkdir -p ./database/{migrations,seeds}

DB_URL=$(cat ./.env | grep DB_URL | cut -d "\"" -f 2)

# if .env not found exit
if [ ! -e ".env" ]; then
  echo "[ERROR:] File \".env\" not found. Create one from .env.example. Exiting."
  exit 1
fi

# check if file exists
if [ ! -e "database/atlas.hcl" ]; then
  echo "[ INFO:] Creating atlas.hcl env file based on .env"

cat << EOF >> ./database/atlas.hcl
env "local" {
  # the real database
  url = $DB_URL?sslmode=disable
  # a temporary database for atlas compute diff in states
  dev = $(dirname $DB_URL)/dev?sslmode=disable
}
EOF
fi

echo "[ INFO:] Obtaining init.sql from dev database..."

TIMEPREFIX=$(date -u +%Y%m%d%H%M%S)
atlas schema inspect --url $DB_URL?sslmode=disable --exclude atlas_schema_revisions --format '{{ sql . }}' > ./database/migrations/${TIMEPREFIX}_init.sql

echo "[ INFO:] Creating migrations hash file"

atlas migrate hash --dir "file://database/migrations"

echo "[   OK:] Done. You can now commit your db schemas."