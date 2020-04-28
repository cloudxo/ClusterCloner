package gke

import (
	container "cloud.google.com/go/container/apiv1"
	"clusterCloner/clusters/utils"
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	containerpb "google.golang.org/genproto/googleapis/container/v1"
	"log"
)

//ReadClusters Return data on the cluster in JSON form. cliCtx shold provide project and location (zone, where _ means all zones)
func ReadClusters(cliCtx *cli.Context) *containerpb.ListClustersResponse {

	ctx := context.Background()
	c, err := container.NewClusterManagerClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	proj := cliCtx.String("project")
	loc := cliCtx.String("location")
	path := fmt.Sprintf("projects/%s/locations/%s", proj, loc)
	req := &containerpb.ListClustersRequest{Parent: path}
	resp, err := c.ListClusters(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintAsJson(resp)
	return resp
}
