package resources

import "embed"

var (
	//go:embed templates
	TemplatesFS embed.FS

	//go:embed assets
	AssetsFS embed.FS
)
