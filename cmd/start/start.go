package start

import (
	"fmt"

	"github.com/spf13/cobra"
)

// StartCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start command",
	Long: `Start command is the main command to start the fetcher or start
	a trading bot for a particular strategy.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}
