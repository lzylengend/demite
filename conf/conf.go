package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type conf struct {
	Port     string `yam:"port"`
	FilePath string `yam:"filepath"`
	DbPath   string `yam:"dbpath"`
	Docpath  string `yam:"docpath"`
	Ip       string `yam:"ip"`
}

var defaltConf *conf

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

	defaltConf = &conf{
		Port:     c.Port,
		FilePath: c.FilePath,
		DbPath:   c.DbPath,
		Docpath:  c.Docpath,
		Ip:       c.Ip,
	}

	_, err = os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}

	return c, err
}

func GetFilePath() string {
	return defaltConf.FilePath
}

func GetPort() string {
	return defaltConf.Port
}

func GetIp() string {
	return defaltConf.Ip
}
