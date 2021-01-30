#!/bin/sh

SONAR_HOST_URL='http://localhost:5300'
BRANCH=$(git branch --show-current)

if [ -z "$1" ] || [ -z "$2" ]
then
  echo "Usage: ./sonnar_scan.sh PROJECT_KEY PROJECT_TOKEN"
else
  sonar-scanner \
    -Dsonar.projectKey=$1 \
    -Dsonar.sources=. \
    -Dsonar.host.url=$SONAR_HOST_URL \
    -Dsonar.branch.name=$BRANCH \
    -Dsonar.login=$2
fi
