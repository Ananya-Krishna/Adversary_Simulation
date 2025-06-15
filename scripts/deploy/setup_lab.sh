#!/bin/bash

# Setup script for purple team lab environment

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Create necessary directories
echo "Creating lab directories..."
mkdir -p lab/docker/{linux-scripts,zeek-logs,zeek-config,suricata-logs,suricata-rules,c2-logs}

# Copy detection rules
echo "Setting up detection rules..."
cp detection_rules/suricata/*.rules lab/docker/suricata-rules/
cp detection_rules/sigma/*.yml lab/docker/zeek-config/

# Build and start containers
echo "Building and starting containers..."
cd lab/docker
docker-compose build --no-cache
docker-compose up -d

# Wait for services to start
echo "Waiting for services to start..."
sleep 10

# Verify services are running
echo "Verifying services..."
docker-compose ps

# Check C2 server logs
echo "Checking C2 server logs..."
docker-compose logs c2-server

echo "Lab environment setup complete!"
echo "C2 server is running on http://localhost:8080"
echo "Linux target: linux-target"
echo "Zeek logs: lab/docker/zeek-logs"
echo "Suricata logs: lab/docker/suricata-logs" 