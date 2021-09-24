package start

import (
	"github.com/sp98/marketmoz/pkg/trade"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var name string
var interval string

var strategyCmd = &cobra.Command{
	Use:   "strategy",
	Short: "Start the strategy server",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("Starting strategy sever", zap.String("name", name))
		err := trade.Start(name)
		if err != nil {
			Logger.Error("failed to start strategies server", zap.Error(err))
		}
	},
}

func init() {
	StartCmd.AddCommand(strategyCmd)

	strategyCmd.Flags().StringVarP(&name, "name", "n", "", "strategy to start")
	strategyCmd.MarkFlagRequired("name")
	strategyCmd.Flags().StringVarP(&interval, "interval", "i", "", "interval")
	strategyCmd.MarkFlagRequired("interval")
}
