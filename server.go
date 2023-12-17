package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type ServerConfig struct{
    host string
    port int
    serviceFilePath string
}

func RunServer(c *ServerConfig){
    LoadServices(c.serviceFilePath)
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)
    http.HandleFunc("/api/services", ServeSelfhostServices)
    http.HandleFunc("/api/services/", ServeSelfhostService)

    address := fmt.Sprintf("%s:%d", c.host, c.port)
    fmt.Printf("listening %s", address)
    if err := http.ListenAndServe(address, nil); err != nil{
        log.Fatal(err)
        os.Exit(1)
    }
}
