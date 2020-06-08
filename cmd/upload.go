package cmd

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/ultramozg/gget/app"
)

func init() {
	rootCmd.AddCommand(uploadCmd)
}

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file",
	Long:  "upload a file",
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
		if err := app.UploadFile(args[0], args[1]); err != nil {
			fmt.Println("Error: ", err)
		}
	},
}
