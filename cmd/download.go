package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/ultramozg/gget/app"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download file",
	Long:  "gget utility is using for download and upload files via http",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("requires an URL and filename")
		}
		_, err := url.ParseRequestURI(args[0])
		if err != nil {
			return fmt.Errorf("invalid URL specified: %s", args[0])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		app.DownloadFile(args[0], args[1])
	},
}
