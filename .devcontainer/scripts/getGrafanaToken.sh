#!/bin/bash
export GRAFANA_URL="$GRAFANA_HOST:$GRAFANA_PORT/"
# make service account
echo "Generating new Grafana API token..."
guid="$(cat /proc/sys/kernel/random/uuid)"

response=$(curl "${GRAFANA_URL}"api/serviceaccounts -u "admin:admin" -d "{\"name\": \"$guid\", \"role\": \"Editor\", \"isDisabled\": false}" -H "Content-Type: application/json")
id=$(echo "$response" | jq -r '.id')

# create a token and store the response
response=$(curl "${GRAFANA_URL}"api/serviceaccounts/"$id"/tokens -u "admin:admin" -d "{\"name\": \"$guid\"}" -H "Content-Type: application/json")
token=$(echo "$response" | jq -r '.key')
export GRAFANA_API_KEY="$token"
# watch all files in pkg/
./.devcontainer/scripts/watch.sh "pkg/"