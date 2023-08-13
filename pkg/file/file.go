package file

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadFromYAML(path string, target interface{}) error {
	yf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(
		yf,
		target,
	)
}
