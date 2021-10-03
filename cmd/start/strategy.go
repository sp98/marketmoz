package start

import (
	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/trade"
	"github.com/sp98/marketmoz/pkg/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var name string
var interval string

var strategyCmd = &cobra.Command{
	Use:   "strategy",
	Short: "Start the trading strategy",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("Starting strategy", zap.String("name", name), zap.String("interval", interval))
		if !utils.Contains(common.DownsamplePeriods, interval) {
			Logger.Error("invalid strategy interval", zap.String("interval", interval))
			return
		}
		err := trade.Start(name, interval)
		if err != nil {
			Logger.Error("failed to run strategy", zap.String("strategy", name), zap.Error(err))
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
