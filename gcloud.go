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
		// Handle error.
	}
	computeService, err = compute.New(client)
	if err != nil {
		// Handle error.
	}
	return computeService
}


func ReadConfig()(gci GCloudInfo, location string){
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

func Start_server()(result *compute.Operation){
	gci := ReadConfig()
	svc := getClient()
	result, err := svc.Instances.Start(gci.ProjectId, gci.Zone, gci.Instance).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}

func Status_server()(result *compute.Instance){
	gci := ReadConfig()
	svc := getClient()
	result, err := svc.Instances.Get(gci.ProjectId, gci.Zone, gci.Instance).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}

func Stop_server()(result *compute.Operation){
	gci := ReadConfig()
	svc := getClient()
	result, err := svc.Instances.Stop(gci.ProjectId, gci.Zone, gci.Instance).Do()
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println(result.Status)
	return result
}
