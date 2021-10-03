/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package get

import (
	"context"

	"github.com/sp98/marketmoz/pkg/common"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var organization string
var token uint32
var cadence string
var exchange string
var segment string

// ohlcCmd represents the ohlc command
var ohlcCmd = &cobra.Command{
	Use:   "ohlc",
	Short: "Get OHLC data",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		db := influx.NewDB(ctx, organization, common.INFLUXDB_URL, common.INFLUXDB_TOKEN)
		defer db.Client.Close()
		instrument := data.NewInstrument("", "", exchange, "", segment, token)
		//query, _ := trade.GetTestQuery()
		query, err := instrument.GetDSQuery(cadence, common.OHLC_QUERY_ASSET)
		if err != nil {
			return
		}
		data, err := instrument.GetOHLC(db, query)
		if err != nil {
			Logger.Error("failed to get ohlc data", zap.Error(err))
			return
		}
		Logger.Info("OHCL data", zap.Any("data", data))
	},
}

func init() {
	GetCmd.AddCommand(ohlcCmd)

	ohlcCmd.Flags().StringVarP(&organization, "organization", "o", "", "influxdb organization")
	ohlcCmd.MarkFlagRequired("organization")
	ohlcCmd.Flags().Uint32VarP(&token, "token", "t", 0, "instruement token")
	ohlcCmd.MarkFlagRequired("token")
	// TODO: limit the values candence can have. For example, only use 1m, 3m, 5m, 1d, 1w, etc.
	ohlcCmd.Flags().StringVarP(&cadence, "cadence", "c", "", "cadence")
	ohlcCmd.MarkFlagRequired("cadence")
	ohlcCmd.Flags().StringVarP(&exchange, "exchange", "e", "", "exchange")
	ohlcCmd.MarkFlagRequired("exchange")
	ohlcCmd.Flags().StringVarP(&segment, "segment", "s", "", "segment")
	ohlcCmd.MarkFlagRequired("segment")
}
