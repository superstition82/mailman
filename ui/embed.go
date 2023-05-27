// Pakcage ui handles the frontend embedding.
package ui

import (
	"embed"
)

//go:embed dist
var DistDir embed.FS
