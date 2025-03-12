package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/K-Phoen/grabana"
	"github.com/fastleansmart/grabana-devcontainer/pkg/dashboards"
)

func main() {
	shouldBuild := flag.Bool("build", false, "Whether to build all dashboards")
	outDir := flag.String("output", "out", "Directory where to build the dashboards to")
	flag.Parse()
	args := flag.Args()

	dashboards, err := dashboards.InitDashboards()
	if err != nil {
		fmt.Printf("Could not initialize dashboards: %s\n", err)
		os.Exit(1)
	}

	dashboardUIDs := make([]string, len(dashboards))
	for i, d := range dashboards {
		dashboardUIDs[i] = d.Internal().UID
	}

	fmt.Printf("Got dashboards: %v\n", dashboardUIDs)

	// only build dashboards as JSON
	if *shouldBuild {
		fmt.Println("Building dashboards...")
		for _, builder := range dashboards {
			bytes, err := builder.MarshalIndentJSON()
			if err != nil {
				fmt.Printf("Could not build dashboard: %s\n", err)
				os.Exit(1)
			}
			path := path.Join(*outDir, normalizeUid(builder.Internal().UID)+".json")
			err = os.WriteFile(path, bytes, 0644)
			if err != nil {
				fmt.Printf("Could not write dashboard to file: %s\n", err)
				os.Exit(1)
			}
		}
		os.Exit(0)
	}

	if len(args) != 2 {
		fmt.Fprint(os.Stderr, "Usage: go run main.go http://grafana-host:3000 api-key-string-here\n")
		os.Exit(1)
	}

	host := strings.TrimRight(args[0], "/")
	ctx := context.Background()
	client := grabana.NewClient(&http.Client{}, host, grabana.WithAPIToken(args[1]))

	folder, err := client.FindOrCreateFolder(ctx, "Grabana generated Dashboards")
	if err != nil {
		fmt.Printf("Could not create folder: %s\n", err)
		os.Exit(1)
	}

	for _, builder := range dashboards {
		dash, err := client.UpsertDashboard(ctx, folder, builder)
		if err != nil {
			fmt.Printf("Could not create dashboard: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("Uploaded dashboard to:\n%s\n", host+dash.URL)
	}
}

func normalizeUid(uid string) string {
	r := regexp.MustCompile(`[^a-z0-9.-]`)
	return r.ReplaceAllString(uid, "-")
}
