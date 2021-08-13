package cmd

import (
	"github.com/sp98/marketmoz/pkg/fetcher"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var source string
var destination string

var Logger *zap.Logger

// fetcherCmd represents the fetcher command
var fetcherCmd = &cobra.Command{
	Use:   "fetcher",
	Short: "Fetcher gets the tick data for a stock",
	Long: `Fetcher gets the tick data for a stock
	and stores it in a database for further analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Info("fetcher called with args", zap.Strings("args", args))
		fetcher.StartFetcher(source, destination)
	},
}

func init() {
	startCmd.AddCommand(fetcherCmd)

	// Set flags
	fetcherCmd.Flags().StringVarP(&source, "source", "s", "", "data source to fetch tick data")
	fetcherCmd.MarkFlagRequired("source")
	fetcherCmd.Flags().StringVarP(&destination, "destination", "d", "", "database to store the tick data")
	fetcherCmd.MarkFlagRequired("destination")
}
