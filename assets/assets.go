package assets

import "embed"

//go:embed *

var f embed.FS

func ReadFile(name string) ([]byte, error) {
	return f.ReadFile(name)
}
