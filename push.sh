#!/bin/bash
msg=${1:-"update"}
timestamp=$(date "+%Y-%m-%d %H:%M:%S")
echo "[$timestamp] %msg" >> .gitpush.log
git add .
git commit -m "$msg"
git push origin main
