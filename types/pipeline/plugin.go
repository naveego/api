package pipeline

// PluginSelector identifies a version or version range of a plugin.
type PluginSelector struct {
	ID      string `json:"id"`      // The ID of the plugin
	Version string `json:"version"` // The version of the plugin implementation.
}

type Plugin struct {
	ID          string          `json:"id,omitempty" bson:"_id,omitempty"` // The ID of the plugin
	Description string          `json:"description,omitempty"`
	Versions    []PluginVersion `json:"versions,omitempty"`
}

type PluginVersion struct {
	Description  string `json:"description,omitempty"`
	Version      string `json:"version,omitempty"`
	ConfigSchema string `json:"configSchema,omitempty"`
	URL          string `json:"url,omitempty"`       // the location the plugin can be downloaded from
	Signature    string `json:"signature,omitempty"` // The SHA-1 hash of the plugin, used to verify download.
}
