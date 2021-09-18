package strategies

import (
	"fmt"
	"net/http"

	"github.com/sp98/marketmoz/pkg/strategies/pvt"
)

const (
	PVT_STRATEGY = "pvt"
)

func StartServer(name string) error {
	// Add handlers
	switch name {
	case PVT_STRATEGY:
		pvt.StrategyExample4()
		//http.HandleFunc(fmt.Sprintf("/%s", name), pvt.Handler)
	default:
		return fmt.Errorf("invalid strategy %q", name)
	}

	// Start server
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		return fmt.Errorf("failed to start strategies server. Error %v", err)
	}

	return nil
}
