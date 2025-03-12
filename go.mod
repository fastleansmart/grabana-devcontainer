module github.com/fastleansmart/grabana-devcontainer

go 1.23.0

replace github.com/K-Phoen/grabana v0.22.2 => github.com/fastleansmart/grabana v0.23.0

replace github.com/K-Phoen/sdk v0.12.4 => github.com/fastleansmart/grafana-sdk v0.14.3

require (
	github.com/K-Phoen/grabana v0.22.2
	github.com/K-Phoen/sdk v0.12.4
)

require (
	github.com/gosimple/slug v1.13.1 // indirect
	github.com/gosimple/unidecode v1.0.1 // indirect
	github.com/prometheus/common v0.45.0 // indirect
)
