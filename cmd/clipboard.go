package cmd

import (
	"io"
	"net/http"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clipboardCmd)
}

var clipboardCmd = &cobra.Command{
	Use:   "clipboard",
	Short: "Share Clipboard content",
	Run: func(cmd *cobra.Command, args []string) {
		data, err := clipboard.ReadAll()
		if err != nil {
			ExitWithError("Can not read clipboard content")
		}
		if len(data) == 0 {
			ExitWithError("Empty clipboard content")
		}

		http.HandleFunc("/", basicAuth(func(w http.ResponseWriter, req *http.Request) {
			io.WriteString(w, data)
		}))

		printQR(cmd)
		http.ListenAndServe(port, nil)
	},
}
