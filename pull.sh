#!/bin/bash
timestamp=$(date "+%Y-%m-%d %H-%M-%S")
echo "[timestamp] Pulling latest changes..." >> .gitpush.log
git pull origin main
