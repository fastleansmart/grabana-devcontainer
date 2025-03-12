#!/usr/bin/env bash
set -eu -o pipefail

echo "Watching files in ${1} directory..."

error() {
	>&2 echo "ERROR:" "${@}"
	exit 1
}

[ -n "${GRAFANA_API_KEY:-}" ] || error "Invalid GRAFANA_API_KEY"
[[ "${GRAFANA_URL:-}" =~ ^https?://[^/]+/$ ]] || error "Invalid GRAFANA_URL (example: 'http://localhost:3001/' incl. slash at end)"

dir_to_watch="${1}"

# watch all files in the given directory
while true; do
	# need to add "|| true", so we ignore errors, because entr throws an error on directory changes
	find "$dir_to_watch" | entr -d -n go run main.go "$GRAFANA_URL" "$GRAFANA_API_KEY" || true
done