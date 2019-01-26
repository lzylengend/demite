package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type conf struct {
	Port string `yam:port`
}

func Init(path string) (*conf, error) {
	c := &conf{}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	return c, err
}
