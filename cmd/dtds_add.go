package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	dtdsCmd.AddCommand(dtdsAddCmd)
}

var dtdsAddCmd = &cobra.Command{
	Use:   "add <path>",
	Short: "Add a path to disable type detection",
	Long:  `Add a path to disable type detection.`,
	Args:  cobra.ExactArgs(1), //nolint:mnd
	Run: python(func(cmd *cobra.Command, args []string, d pythonData) {
		s, err := d.store.Settings.Get()
		checkErr(err)
		s.DisableTypeDetections = append(s.DisableTypeDetections, args[0])
		err = d.store.Settings.Save(s)
		checkErr(err)
		printDtds(s.DisableTypeDetections)
	}, pythonConfig{}),
}
