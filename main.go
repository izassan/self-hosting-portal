package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
    HOST = "host"
    PORT = "port"
    SERVICEFILEPATH = "service-filepath"
)

var rootCmd = &cobra.Command{
    Use: "self-hosting portal",
    RunE: func(cmd *cobra.Command, args []string) error{
        host, err := cmd.Flags().GetString(HOST)
        if err != nil{
            return err
        }
        port, err := cmd.Flags().GetInt(PORT)
        if err != nil{
            return err
        }
        serviceFilePath, err := cmd.Flags().GetString(SERVICEFILEPATH)
        if err != nil{
            return err
        }
        RunServer(&ServerConfig{
            host: host,
            port: port,
            serviceFilePath: serviceFilePath,
        })
        return nil
    },
}

func main(){
    if err := rootCmd.Execute(); err != nil{
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func init(){
    rootCmd.Flags().StringP(HOST, "H", "", "server host")
    rootCmd.Flags().IntP(PORT, "p", 7426, "server port")
    rootCmd.Flags().StringP(SERVICEFILEPATH, "f", "./services.json", "service definition file path")
}
