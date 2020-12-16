package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	dtdsCmd.AddCommand(dtdsRmCmd)
}

var dtdsRmCmd = &cobra.Command{
	Use:   "rm <index> [index_end]",
	Short: "Removes a path to disable type detection",
	Long: `Removes a path to disable type detection. The provided index
is the same that's printed when you run 'dtds ls'. Note
that after each removal/addition, the index of the
paths change. So be careful when removing them after each
other.

You can also specify an optional parameter (index_end) so
you can remove all paths from 'index' to 'index_end',
including 'index_end'.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.RangeArgs(1, 2)(cmd, args); err != nil { //nolint:mnd
			return err
		}

		for _, arg := range args {
			if _, err := strconv.Atoi(arg); err != nil {
				return err
			}
		}

		return nil
	},
	Run: python(func(cmd *cobra.Command, args []string, d pythonData) {
		s, err := d.store.Settings.Get()
		checkErr(err)

		i, err := strconv.Atoi(args[0])
		checkErr(err)
		f := i
		if len(args) == 2 { //nolint:mnd
			f, err = strconv.Atoi(args[1])
			checkErr(err)
		}

		s.DisableTypeDetections = append(s.DisableTypeDetections[:i], s.DisableTypeDetections[f+1:]...)
		err = d.store.Settings.Save(s)
		checkErr(err)
		printDtds(s.DisableTypeDetections)
	}, pythonConfig{}),
}
