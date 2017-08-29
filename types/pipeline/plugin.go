package pipeline

// PluginVersion represents a version of a publisher or subscriber plugin.
type PluginVersion struct {
	ID      string `json:"id"`      // The ID of the plugin
	Version string `json:"version"` // The version of the plugin type
}
