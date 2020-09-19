package conf

import "encoding/json"

type Config struct {
	ExcludeDirectories []string `json:"excludeDirectories,omitempty"`
	ExcludeFormats     []string `sjson:"excludeFormats,omitempty"`
}

func (c *Config) UnmarshalJSON(j []byte) (*Config, error) {
	var genConfig Config
	err := json.Unmarshal(j, &genConfig)
	if err != nil {
		panic(err)
	}
	return &genConfig, nil
}
