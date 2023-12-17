package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

type SelfhostServices struct{
    Services []SelfhostService `json:"services"`
}

type SelfhostService struct{
    Name string `json:"name"`
    Description string `json:"description"`
    ServiceURL string `json:"service_url"`
    ImageURL string `json:"img_url"`
}

var services SelfhostServices

func LoadServices(fp string) error{
    raw, err := ioutil.ReadFile(fp)
    if err != nil{
        return err
    }
    json.Unmarshal(raw, &services)
    return nil
}

func ServeSelfhostServices(w http.ResponseWriter, r *http.Request){
    var responseServiceNames struct{
        Services []string `json:"services"`
    }
    serviceNames := []string{}
    for _, service := range services.Services{
        serviceNames = append(serviceNames, service.Name)
    }

    responseServiceNames.Services = serviceNames

    j, err := json.Marshal(responseServiceNames)
    if err != nil{
        log.Printf("%v", err)
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintf(w, "failed marshal")
        return
    }
    fmt.Fprintf(w, string(j))
}

func ServeSelfhostService(w http.ResponseWriter, r *http.Request){
    sub := strings.TrimPrefix(r.URL.Path, "/api/services")
    _, id := filepath.Split(sub)
    found := false
    var foundService SelfhostService
    for _, service := range services.Services{
        if id == service.Name{
            found = true
            foundService = service
        }
    }

    if found {
        j, err := json.Marshal(foundService)
        if err != nil{
            log.Printf("%v", err)
            w.WriteHeader(http.StatusInternalServerError)
            fmt.Fprintf(w, "failed marshal")
            return
            }
        fmt.Fprintf(w, string(j))
        return
    }
    w.WriteHeader(http.StatusNotFound)
    fmt.Fprintf(w, "not found service")
}
