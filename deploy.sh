#!/bin/bash

# --- CONFIG ---
KEY_PATH="~/.ssh/lightsail_key.pem"
USER="ubuntu"
HOST="34.192.2.8"
SERVICE_NAME="fdecono"
LOCAL_BINARY_NAME="fdecono"
LOCAL_BINARY_PATH="./fdecono"
LOCAL_ENTRY_POINT="cmd/server/main.go"
REMOTE_BINARY="/usr/local/bin/fdecono"
TMP_PATH="/tmp/fdecono"

# --- BUILD ---
echo "Building Go binary for Linux..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $LOCAL_BINARY_NAME $LOCAL_ENTRY_POINT || { echo "Build failed"; exit 1; }

# --- UPLOAD ---
echo "Uploading binary to Lightsail..."
scp -i $KEY_PATH $LOCAL_BINARY_PATH $USER@$HOST:$TMP_PATH || { echo "Upload failed"; exit 1; }

# --- REPLACE & RESTART ---
echo "Replacing binary and restarting service..."
ssh -i $KEY_PATH $USER@$HOST << EOF
sudo systemctl stop $SERVICE_NAME
sudo mv $TMP_PATH $REMOTE_BINARY
sudo chown fdecono-site:fdecono-site $REMOTE_BINARY
sudo chmod +x $REMOTE_BINARY
sudo systemctl daemon-reload
sudo systemctl start $SERVICE_NAME
sudo systemctl status $SERVICE_NAME -l --no-pager
EOF

echo "Deployment finished!"

