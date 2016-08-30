package types

import (
	"strings"
)

// Repository represents the core multi-tenancy boundary in Naveego
type Repository struct {
	ID     string                 `json:"id"`     // The id of the repository
	Name   string                 `json:"name"`   // The name of the repository
	GoLive string                 `json:"goLive"` // The date the repository went live
	Config map[string]interface{} `json:"config"` // The configuration of the Repository
}

// HasConfigParam provides a method to determine if a configuration value is defined
// in the repository or not
func (r *Repository) HasConfigParam(paramName string) bool {
	_, exists := getConfigParamRecursive(r.Config, paramName)
	return exists
}

func (r *Repository) GetConfig(paramName string) interface{} {
	val, _ := getConfigParamRecursive(r.Config, paramName)
	return val
}

func (r *Repository) GetConfigWithDefault(paramName string, defaultVal interface{}) interface{} {
	val, exists := getConfigParamRecursive(r.Config, paramName)
	if !exists {
		return defaultVal
	}
	return val
}

func getConfigParamRecursive(config map[string]interface{}, path string) (interface{}, bool) {

	pathParts := strings.Split(path, ".")
	currentProp := pathParts[0]

	val, ok := config[currentProp]
	if !ok {
		return nil, false
	}

	if len(pathParts) == 1 {
		return val, true
	}

	switch v := val.(type) {
	case map[string]interface{}:
		subPath := strings.Join(pathParts[1:], ".")
		return getConfigParamRecursive(v, subPath)
	}

	return nil, false

}
