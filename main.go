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
package main

import (
	"time"

	"github.com/sp98/marketmoz/cmd/delete"
	"github.com/sp98/marketmoz/cmd/get"
	"github.com/sp98/marketmoz/cmd/marketmoz"
	"github.com/sp98/marketmoz/cmd/start"
	"github.com/sp98/marketmoz/pkg/data"
	"github.com/sp98/marketmoz/pkg/db/influx"
	"github.com/sp98/marketmoz/pkg/fetcher"
	"github.com/sp98/marketmoz/pkg/fetcher/kite"
	"github.com/sp98/marketmoz/pkg/trade"
	"github.com/sp98/marketmoz/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	initLogger()
	addCommands()
	marketmoz.Execute()
}

func addCommands() {
	marketmoz.RootCmd.AddCommand(
		start.StartCmd,
		delete.DeleteCmd,
		get.GetCmd,
	)

}
func initLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, _ := config.Build()

	start.Logger = logger
	fetcher.Logger = logger
	utils.Logger = logger
	kite.Logger = logger
	influx.Logger = logger
	data.Logger = logger
	get.Logger = logger
	trade.Logger = logger
}
