#! /usr/bin/env bash

NEW_STATUS=$1

if [[ -z "$NEW_STATUS" ]]; then
  echo "Usage: $0 <new status>"
  exit 1
fi

if [[ -z "$PROGRAMMERBAR_API_TOKEN" ]]; then
  echo "PROGRAMMERBAR_API_TOKEN is not set"
  exit 1
fi

curl -X POST https://programmerbar.fly.dev/status \
    -H "Authorization: Bearer $PROGRAMMERBAR_API_TOKEN" \
    -d "{\"status\": $NEW_STATUS}"
