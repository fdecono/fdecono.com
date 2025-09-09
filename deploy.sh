#!/bin/bash

# Simple deployment script for fdecono.com

echo "🚀 Deploying fdecono.com..."

# Build the application
echo "📦 Building application..."
go build -o fdecono cmd/server/main.go

if [ $? -ne 0 ]; then
    echo "❌ Build failed!"
    exit 1
fi

echo "✅ Build successful!"

# Check if we're deploying to production
if [ "$1" = "production" ]; then
    echo "🌐 Deploying to production..."
    
    # Add your production deployment commands here
    # For example, if using a VPS:
    # scp fdecono user@your-server:/home/user/
    # scp -r internal/templates user@your-server:/home/user/
    # scp -r internal/static user@your-server:/home/user/
    # ssh user@your-server "sudo systemctl restart fdecono"
    
    echo "⚠️  Production deployment not configured yet."
    echo "   Please update deploy.sh with your production server details."
else
    echo "🏃 Running locally..."
    ./fdecono
fi
