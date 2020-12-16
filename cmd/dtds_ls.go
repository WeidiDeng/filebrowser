package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	dtdsCmd.AddCommand(dtdsLsCmd)
}

var dtdsLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all paths to disable type detection",
	Long:  `List all paths to disable type detection.`,
	Args:  cobra.NoArgs,
	Run: python(func(cmd *cobra.Command, args []string, d pythonData) {
		s, err := d.store.Settings.Get()
		checkErr(err)
		printDtds(s.DisableTypeDetections)
	}, pythonConfig{}),
}
