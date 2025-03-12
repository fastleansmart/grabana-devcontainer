#!/bin/bash

# check if grafana is up
echo "Checking if Grafana is running..."
curl -s -f -o /dev/null host.docker.internal:"$GRAFANA_PORT" || { >&2 printf '\nERROR: Grafana not found. Have you set GRAFANA_PORT correctly?\n'; exit 1; }

# install some libraries
echo "Installing libraries..."
go mod tidy

# install entr for file watching
git clone https://github.com/eradman/entr.git /tmp/entr
cd /tmp/entr || exit
./configure
sudo make install