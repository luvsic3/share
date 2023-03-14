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

var IpFlag string

type any = interface{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "share /path/to/directory",
	Short: "Share directories and files",
	Args:  cobra.MinimumNArgs(1),
	Long:  `Share directories and files from the CLI to iOS and Android devices without the need of an extra client app`,
	Run: func(cmd *cobra.Command, args []string) {
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

		printQR(cmd)

		http.ListenAndServe(port, http.FileServer(http.Dir(targetDir)))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().StringVar(&IpFlag, "ip", "", "Your machine public ip address")
	cobra.CheckErr(rootCmd.Execute())
}

func ExitWithError(v any) {
	fmt.Println(v)
	os.Exit(1)
}

func printQR(cmd *cobra.Command) {
	cmd.Println("Scan the QR-Code to access directory on your phone")
	cmd.Println()

	var ip string = IpFlag
	if IpFlag == "" {
		ip = GetOutboundIP().String()
	}

	url := protocol + ip + port

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
