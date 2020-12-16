package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dtdsCmd)
}

var dtdsCmd = &cobra.Command{
	Use:   "dtds",
	Short: "Disable type detection management utility",
	Long:  `Disable type detection management utility.`,
	Args:  cobra.NoArgs,
}

func printDtds(dtds []string) {
	for i, dtd := range dtds {
		fmt.Printf("(%d): %s\n", i, dtd)
	}
}
