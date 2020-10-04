package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ncrypt96/gena/constants"
)

type Config struct {
	Variables map[string]string `json:"variables"`
	AddGit    bool              `json:"addGit"`
	Postgen   string            `json:"postgen"`
}

func NewConfig() *Config {
	return new(Config)
}

func (c *Config) UnmarshalConfig(j []byte) error {
	return json.Unmarshal(j, &c)
}

func (c *Config) GetConfigAsBytes(f string) ([]byte, error) {
	cf, err := os.Open(fmt.Sprintf("./%s/%s", f, constants.ConfFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("The repository is not a valid Gena repository \n \t The repository must have a valid %s at the root of the repository", constants.ConfFileName)
		}
		return nil, err
	}
	b, err := ioutil.ReadAll(cf)
	if err != nil {
		return nil, err
	}
	return b, nil
}
