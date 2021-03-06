package integrationtests

import (
	"clustercloner/clusters"
	"testing"
)

func TestCreateGCPClusterFromFileThenCloneToAzure(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	file := "test-data/gke_clusters.json"
	outputCloud := clusters.Azure
	runClusterCloning(t, file, outputCloud)
}
