package mandoscontroller

import mjparse "github.com/Reshusk23/dme-util/test-util/mandos/json/parse"

// NewDefaultFileResolver yields a new DefaultFileResolver instance.
// Reexported here to avoid having all external packages importing the parser.
// DefaultFileResolver is in parse for local tests only.
func NewDefaultFileResolver() *mjparse.DefaultFileResolver {
	return mjparse.NewDefaultFileResolver()
}
