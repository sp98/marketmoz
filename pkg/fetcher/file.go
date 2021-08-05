package fetcher

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sp98/marketmoz/assets"
)

const (
	niftyAsset = "data/nifty.txt"
)

type FileData struct {
	Stock string
	Open  float64
	High  float64
	Low   float64
	Close float64
}

// startFileFetcher starts the fetching process from the file
func startFileFetcher() error {
	dataBytes, err := assets.ReadFile(niftyAsset)
	if err != nil {
		return fmt.Errorf("failed to read file %q", niftyAsset)
	}

	dataString := string(dataBytes)
	dataList := []FileData{}
	lines := strings.Split(dataString, "\n")
	for _, line := range lines {
		l := strings.Split(line, ",")
		if len(l) > 6 {
			open, _ := strconv.ParseFloat(l[3], 64)
			high, _ := strconv.ParseFloat(l[4], 64)
			low, _ := strconv.ParseFloat(l[5], 64)
			close, _ := strconv.ParseFloat(l[6], 64)

			dataList = append(dataList, FileData{
				Stock: l[0],
				Open:  open,
				High:  high,
				Low:   low,
				Close: close,
			})
		}
	}

	fmt.Println("total data - ", dataList)

	return nil
}
