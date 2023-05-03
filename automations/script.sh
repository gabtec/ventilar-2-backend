#! /bin/bash

echo "Checking if database schemas have changes"

LOCALDIR=$(dirname $0)

db_url=$(cat ./backend/.env | grep DB_STRING | cut -d "\"" -f 2)

echo $db_url

status=$(shasum -a 256 ./backend/\*.model.go -c ./backend/automations/models-sums.txt)

echo $status