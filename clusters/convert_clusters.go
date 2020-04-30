package clusters

import (
	"errors"
	"fmt"
	"log"
)

var (
	HUB   = "Hub"
	GCP   = "GCP"
	AZURE = "Azure"
	AWS   = "AWS"
)

func Transform(clusterInfo ClusterInfo, toCloud string) (c ClusterInfo, err error) {
	hub, err1 := ToHubFormat(clusterInfo)
	if err1 != nil {
		return ClusterInfo{}, err1
	}
	target, err2 := FromHubFormat(hub, toCloud)
	if err2 != nil {
		return ClusterInfo{}, err2
	}
	return target, nil
}
func ToHubFormat(clusterInfo ClusterInfo) (c ClusterInfo, err error) {
	err = nil
	var ret ClusterInfo
	switch cloud := clusterInfo.Cloud; cloud { //todo Split out into "adapters" to avoid switch statement. Putting it here now for reference.
	case GCP:
		log.Print("From GCP ")
		ret, err = ConvertGCPToHub(clusterInfo)
	case AZURE:
		log.Print("From Azure ")
		ret, err = ConvertAzureToHub(clusterInfo)
	case AWS:
		log.Print("From AWS ")
		return c, errors.New(fmt.Sprintf("Unsupported %s", cloud))
	case HUB:
		log.Print("From Hub , no changes")
		ret = c
	default:
		return c, errors.New(fmt.Sprintf("Unknown %s", cloud))
	}
	return ret, err
}
func FromHubFormat(clusterInfo ClusterInfo, toCloud string) (c ClusterInfo, err error) {
	if clusterInfo.Cloud != HUB {
		return ClusterInfo{}, errors.New(fmt.Sprintf("Wrong Cloud %s", clusterInfo.Cloud))
	}
	err = nil
	var ret ClusterInfo
	switch toCloud { //todo Split out into "adapters" to avoid switch statement. Putting it here now for reference.
	case GCP:
		log.Print("to GCP ")
		ret, err = ConvertHubToGCP(clusterInfo)
	case AZURE:
		log.Print("to Azure ")
		ret, err = ConvertHubToAzure(clusterInfo)
	case AWS:
		log.Print("to AWS ")
		return c, errors.New(fmt.Sprintf("Unsupported %s", toCloud))
	case HUB:
		log.Print("to Hub , no changes")
		ret = c
	default:
		return c, errors.New(fmt.Sprintf("Unknown %s", toCloud))
	}
	return ret, err
}

func ConvertAzureToHub(clusterInfo ClusterInfo) (ClusterInfo, error) {
	var ret = clusterInfo
	ret.sourceCluster = &clusterInfo
	if clusterInfo.sourceCluster == ret.sourceCluster {
		panic("Copying didn't work as expected")
	}
	ret.Cloud = HUB
	// ret.Name unchanged
	// ret.NodeCount unchanged
	ret.Scope = "" //Scope not meaningful in conversion cross-cloud
	loc, err := ConvertLocationAzureToHub(ret.Location)
	ret.Location = loc
	return ret, err
}

func ConvertHubToAzure(clusterInfo ClusterInfo) (ClusterInfo, error) {
	//todo this is duplicate to ConvertAzureToHub
	var ret = clusterInfo
	ret.sourceCluster = &clusterInfo
	if clusterInfo.sourceCluster == ret.sourceCluster {
		panic("Copying didn't work as expected")
	}
	ret.Cloud = AZURE
	// ret.Name unchanged
	// ret.NodeCount unchanged
	ret.Scope = "" //Scope not meaningful in conversion cross-cloud
	loc, err := ConvertLocationHubToAzure(ret.Location)
	ret.Location = loc
	return ret, err
}

func ConvertGCPToHub(clusterInfo ClusterInfo) (ClusterInfo, error) {
	//todo this is duplicate to ConvertAzureToHub
	var ret = clusterInfo
	ret.sourceCluster = &clusterInfo
	if clusterInfo.sourceCluster == ret.sourceCluster {
		panic("Copying didn't work as expected")
	}
	ret.Cloud = HUB
	//	ret.Name unchanged
	// ret.NodeCount unchanged
	ret.Scope = "" //Scope not meaningful in converstion cross-cloud
	loc, err := ConvertLocationGcpToHub(ret.Location)
	ret.Location = loc
	return ret, err
}

func ConvertHubToGCP(clusterInfo ClusterInfo) (ClusterInfo, error) {
	//todo this is duplicate to ConvertAzureToHub
	var ret = clusterInfo
	ret.sourceCluster = &clusterInfo
	if clusterInfo.sourceCluster == ret.sourceCluster {
		panic("Copying didn't work as expected")
	}
	ret.Cloud = GCP
	//	ret.Name unchanged
	// ret.NodeCount unchanged
	ret.Scope = "" //Scope not meaningful in conversion cross-cloud
	loc, err := ConvertLocationHubToToGcp(ret.Location)
	ret.Location = loc
	return ret, err
}