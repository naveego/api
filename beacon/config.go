package beacon

import (
	"encoding/json"
	"strings"
)

type Config struct {
	data map[string]interface{}
}

// UnmarshalJSON provides custom deserialization for config.
func (c *Config) UnmarshalJSON(b []byte) error {
	c.data = make(map[string]interface{})
	return json.Unmarshal(b, &c.data)
}

// HasConfigParam provides a method to determine if a configuration value is defined
// in the Config or not.
func (c *Config) HasConfigParam(paramName string) bool {
	_, exists := getConfigParamRecursive(c.data, paramName)
	return exists
}

func (c *Config) GetConfig(paramName string) interface{} {
	val, _ := getConfigParamRecursive(c.data, paramName)
	return val
}

func (c *Config) GetConfigWithDefault(paramName string, defaultVal interface{}) interface{} {
	val, exists := getConfigParamRecursive(c.data, paramName)
	if !exists {
		return defaultVal
	}
	return val
}

func (c *Config) GetStringConfig(paramName string) (string, bool) {
	raw := c.GetConfig(paramName)
	if raw != nil {
		s, ok := raw.(string)
		if ok {
			return s, true
		}
	}

	return "", false
}

func getConfigParamRecursive(config map[string]interface{}, path string) (interface{}, bool) {

	index := strings.Index(path, ".")
	var currentProp string
	if index < 0 {
		currentProp = path
	} else {
		currentProp = path[:index]
	}

	val, ok := config[currentProp]
	if !ok {
		return nil, false
	}

	if len(path) == len(currentProp) {
		return val, true
	}

	switch v := val.(type) {
	case map[string]interface{}:
		subPath := path[index+1:]
		return getConfigParamRecursive(v, subPath)
	}

	return nil, false

}
