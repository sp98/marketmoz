package cmd

import (
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var organization string
var bucket string

// dataCmd represents the data command
var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Data represents the influx data in the application",
	Long: `
Data represents the influx data in the application`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := data.Delete(organization, bucket); err != nil {
			Logger.Fatal("failed to delete data", zap.Error(err))
		}
	},
}

func init() {
	deleteCmd.AddCommand(dataCmd)
	dataCmd.Flags().StringVarP(&organization, "organization", "o", "", "influxdb organization")
	dataCmd.MarkFlagRequired("organization")
	dataCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "influxdb bucket")
	dataCmd.MarkFlagRequired("bucket")
}
