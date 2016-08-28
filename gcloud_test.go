package gcloud

import (
	"testing"
	"log"
)

var(
	Gci   GCloudInfo
)

func TestStartServer(t *testing.T){
	Gci = ReadConfig("./gcloud_config.yml")

	Start_server(Gci)

}

func TestReadConfig(t *testing.T) {
	if Gci.Instance != "ftb-infinity-server-2"{
		log.Fatalf("unexpected gci.Instance, got: %v", Gci.Instance)
	}
	if Gci.Zone != "us-central1-a"{
		log.Fatalf("unexpected gci.Zone, got: %v", Gci.Zone)
	}
	if Gci.ProjectId != "silent-space-421"{
		log.Fatalf("unexpected gci.ProjectId, got: %v", Gci.ProjectId)
	}
}