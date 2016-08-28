package gcloud

import (
	"testing"
	"log"
)

func TestStartServer(t *testing.T){
	Start_server()

}

func TestReadConfig(t *testing.T) {
	gci := ReadConfig()
	if gci.Instance != "ftb-infinity-server-2"{
		log.Fatalf("unexpected gci.Instance, got: %v", gci.Instance)
	}
	if gci.Zone != "us-central1-a"{
		log.Fatalf("unexpected gci.Zone, got: %v", gci.Zone)
	}
	if gci.ProjectId != "silent-space-421"{
		log.Fatalf("unexpected gci.ProjectId, got: %v", gci.ProjectId)
	}
}