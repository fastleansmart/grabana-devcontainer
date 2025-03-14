# Grabana Devcontainer

This repo provides an easy way to develop [Grafana](https://grafana.com/) dashboards using [Grabana](https://github.com/K-Phoen/grabana/).

With this, you'll able to code your dashboards and see the changes in real time in your local Grafana.

## Getting started

Before opening this repo, make sure that your target Grafana is already running. You could for example use the [grafana docker container](https://hub.docker.com/r/grafana/grafana).

The port of your instance should be set in `$GRAFANA_PORT`.
The host URL should be in `$GRAFANA_HOST`.

Next, you need to open this repository in the existing [devcontainer](https://code.visualstudio.com/docs/devcontainers/tutorial) in VSCode.

After the container has finished building, you are ready to start building your beautiful dashboards.

Dashboards are generally defined under `pkg/dashboards/dashboard/yourAmazingDashboard.go`. Be sure to register your dashboard in `pkg/dashboards/dashboards.go` for them to be built.

## Publishing dashboards

To publish dashboards to your real Grafana instance, build them using the following command:

```bash
go run main.go --build --output <output-dir>
```

Alternatively, if you want to publish your dashboards using the Grafana API, simply run:

```bash
go run main.go <your-grafana-url> <your-grafana-api-token>
```

Afterwards, just simply upload the JSON output into your Grafana and you're done!
