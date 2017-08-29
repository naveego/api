package pipeline

import (
	"github.com/blang/semver"
)

// PluginVersion represents a version of a publisher or subscriber plugin.
type PluginVersion struct {
	ID      string         `json:"id"`      // The ID of the node
	Version semver.Version `json:"version"` // The nodes type
}
