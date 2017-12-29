package cmd

import "github.com/spf13/cobra"

type Options struct {
	optint int
}


var o = &Options{}


var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Short help for show",
	Long:  "Long help for show",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("show called")
	// },
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Printf("show called: optint: %d", o.optint)
		// i, err := cmd.Flags().GetBool("boolean")
		// if err != nil  {
		// 	return err
		// }
		// s, err := cmd.Flags.GetInt("integer")
		// if err != nil  {
		// 	return err
		// }
		// fmt.Println("show called")
		return nil
	},
}

func init() {
	RootCmd.AddCommand(showCmd)

	showCmd.Flags().IntVarP(&o.optint, "int", "i", 0, "int option")
}
