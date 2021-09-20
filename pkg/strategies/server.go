package strategies

import (
	"fmt"
)

const (
	PVT_STRATEGY = "pvt"
)

func Start(name string) error {
	switch name {
	case PVT_STRATEGY:
		exampleStrategy()
	default:
		return fmt.Errorf("invalid strategy %q", name)
	}
	return nil
}
