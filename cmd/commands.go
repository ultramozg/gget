package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "download",
	Short: "Download file",
	Long:  "gget utility is using for download and upload files via http",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires a URL")
		}
		/*if myapp.IsValidColor(args[0]) {
			return nil
		}*/
		return fmt.Errorf("invalid URL specified: %s", args[1])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
