/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"

	"github.com/mdp/qrterminal/v3"
	"github.com/spf13/cobra"
)

const port = ":8080"
const protocol = "http://"

type any = interface{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-share",
	Short: "Share files",
	Long:  `Share files.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			ExitWithError("Must provide directory path")
		}

		var argDir = args[0]
		workDir, err := os.Getwd()

		if err != nil {
			ExitWithError(err)
		}

		var targetDir = path.Join(workDir, argDir)
		info, err := os.Stat(targetDir)
		if err != nil {
			ExitWithError(err)
		}

		if !info.IsDir() {
			ExitWithError("Given path is not a directory")
		}
		cmd.Println("Scan the QR-Code to access:", targetDir, "directory on your phone")
		cmd.Println()

		ip := GetOutboundIP()
		url := protocol + ip.String() + port

		qrterminal.GenerateWithConfig(url, qrterminal.Config{
			Writer:    os.Stdout,
			Level:     qrterminal.L,
			BlackChar: qrterminal.BLACK,
			WhiteChar: qrterminal.WHITE,
			QuietZone: 1,
		})

		cmd.Println()
		cmd.Println("Or access this link: ", url)
		cmd.Println()
		cmd.Println("Press ctrl+c to stop sharing")

		http.ListenAndServe(port, http.FileServer(http.Dir(targetDir)))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func ExitWithError(v any) {
	fmt.Println(v)
	os.Exit(1)
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
