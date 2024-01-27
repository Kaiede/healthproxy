#!/bin/sh

if [ "$ADDRESS" == "" ]; then
  echo "Need to set ADDRESS to be checked."
  exit 1
fi

echo "Starting Healthcheck Proxy"
exec /healthproxy/healthproxy "$ADDRESS"