package gcloud

import (
	"path/filepath"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"google.golang.org/api/compute/v1"
	"golang.org/x/oauth2/google"
	"golang.org/x/net/context"
)

type GCloudInfo struct {
	ProjectId string
	Zone string
	Instance string
}

func getClient()(computeService *compute.Service) {
	ctx := context.TODO()

	client, err := google.DefaultClient(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	computeService, err = compute.New(client)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return computeService
}



func ReadConfig(location string)(gci GCloudInfo){
	filename, _ := filepath.Abs(location)
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	gci = GCloudInfo{}

	err = yaml.Unmarshal(yamlFile, &gci)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return gci
}

func Start_server(gci GCloudInfo, InstanceName string)(result *compute.Operation){
	svc := getClient()
	result, err := svc.Instances.Start(gci.ProjectId, gci.Zone, InstanceName).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}


func New_server(gci GCloudInfo, name string)(result *compute.Operation){
	svc := getClient()

	//TODO: switch to https://github.com/google/google-api-go-client
	new_instance := &compute.Instance{
		Name: name,
		MachineType: "zones/us-central1-a/machineTypes/n1-standard-1",
		Disks: []*compute.AttachedDisk{
			{
				AutoDelete: true,
				Boot:       true,
				Type:       "PERSISTENT",
				InitializeParams: &compute.AttachedDiskInitializeParams{
					DiskName:    name + "-boot",
					SourceImage: "/projects/coreos-cloud/global/images/coreos-stable-1122-2-0-v20160906",
				},
			},
		},
		NetworkInterfaces: []*compute.NetworkInterface{
			{
				AccessConfigs: []*compute.AccessConfig{
					{
						Type: "ONE_TO_ONE_NAT",
						Name: "External NAT",
					},
				},
				Network: "/global/networks/default",
			},
		},
	}
	result, err := svc.Instances.Insert(gci.ProjectId, gci.Zone, new_instance).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}

func Delete_server(gci GCloudInfo, name string)(result *compute.Operation){
	//TODO: make sure you're not accidentally deleting something important
	svc := getClient()
	log.Println("Deleting server with name: " + name)
	result, err := svc.Instances.Delete(gci.ProjectId, gci.Zone, name).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}

func Status_server(gci GCloudInfo, name string)(result *compute.Instance){
	svc := getClient()
	result, err := svc.Instances.Get(gci.ProjectId, gci.Zone, name).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}

func Stop_server(gci GCloudInfo, name string)(result *compute.Operation){
	svc := getClient()
	result, err := svc.Instances.Stop(gci.ProjectId, gci.Zone, name).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}
