package main

import (
	"log"
	"net/http"
	"./gcloud"

	"github.com/husobee/vestigo"
	"encoding/json"
)

var(
	Gci   gcloud.GCloudInfo
)

func main() {
	Gci = gcloud.ReadConfig("./test_mc-worker_config.yml")
	router := vestigo.NewRouter()

	router.Get("/minecraft", GetMinecraftListHandler)
	router.Get("/minecraft/:server_name", GetMinecraftServerHandler)
	router.Get("/minecraft/:server_name/ip", GetServerIpHandler)
	router.Get("/minecraft/:server_name/status", GetServerStatusHandler)

	router.Put("/minecraft/:server_name", CreateMinecraftServerHandler)

	router.Post("/minecraft/:server_name", UpdateMinecraftServerHandler)

	router.Delete("/minecraft/:server_name", DeleteMinecraftServerHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}


func GetMinecraftListHandler(w http.ResponseWriter, r *http.Request) {
	response_body := ""
	w.WriteHeader(200)
	w.Write([]byte(response_body))
}

func GetMinecraftServerHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	result := gcloud.Status_server(Gci, name)
	response_body, _ := result.MarshalJSON()
	w.WriteHeader(200)
	w.Write([]byte(response_body))
}

func GetServerIpHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	result := gcloud.Status_server(Gci, name)
	response_body := result.NetworkInterfaces[0].AccessConfigs[0].NatIP
	w.WriteHeader(200)
	w.Write([]byte(response_body))
}

func GetServerStatusHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	result := gcloud.Status_server(Gci, name)
	response_body := result.Status
	w.WriteHeader(200)
	w.Write([]byte(response_body))
}

func CreateMinecraftServerHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	gcloud.New_server(Gci, name)
	response_body := ""
	w.WriteHeader(200)
	w.Write([]byte(response_body))

}

type PostInfo struct {
	Power string
}

func UpdateMinecraftServerHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	response_body := ""
	decoder := json.NewDecoder(r.Body)
	var body PostInfo
	err := decoder.Decode(&body)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error: " + err.Error()))
	}
	if body.Power == "on"{
		gcloud.Start_server(Gci, name)
	}
	if body.Power == "off"{
		gcloud.Stop_server(Gci, name)
	}

	w.WriteHeader(200)
	w.Write([]byte(response_body))

}

func DeleteMinecraftServerHandler(w http.ResponseWriter, r *http.Request){
	name := vestigo.Param(r, "server_name")
	response_body := ""
	gcloud.Delete_server(Gci, name)
	w.WriteHeader(200)
	w.Write([]byte(response_body))

}
