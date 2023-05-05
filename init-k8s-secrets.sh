#!/bin/bash

FILENAME="ventilar-secret"

echo "[ INFO:] Creating ./k8s/${FILENAME}.yaml file"

mkdir -p ./k8s

# # Read from user input
# echo "What's the desired database name?"
# read DB_NAME

# echo "What's the desired database user?"
# read DB_USER

# echo "What's the desired database password?"
# read DB_PASSW

# bypass user input
DB_USER=admin
DB_PASSW=admin
DB_NAME=ventilar_db
DB_HOST="10.43.22.86"

DB_URL="postgres://${DB_USER}:${DB_PASSW}@${DB_HOST}:5432/${DB_NAME}?sslmode=disable"

# Build atlas.hcl file
cat << EOF > ./k8s/atlas.hcl
env "k8s" {
  url = "${DB_URL}"
}
EOF

# Build secret.yaml file
cat << EOF > ./k8s/${FILENAME}.yaml
apiVersion: v1
kind: Secret
metadata:
  name: ${FILENAME}
type: Opaque
data:
  # e.g. ${DB_URL}
  DB_URL: $(echo $DB_URL | base64)
  DB_USER: $(echo $DB_USER | base64)
  DB_NAME: $(echo $DB_NAME | base64)
  DB_PASSW: $(echo $DB_PASSW | base64)
  atlas.hcl: $(cat ./k8s/atlas.hcl | base64)
EOF

echo "[   OK:] Done."