// 读取配置文件
// author: baoqiang
// time: 2019/3/21 下午8:05
package helper

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"log"
	"path"
	"path/filepath"
)

type Config struct {
	Yuedu  Yuedu  `yaml:"yuedu"`
	Wukong Wukong `yaml:"wukong"`
}

type Yuedu struct {
	DbFile string `yaml:"db_file"`
	DlPath string `yaml:"dl_path"`
}

type Wukong struct {
	DictFile     string `yaml:"dict_file"`
	IndexPath    string `yaml:"index_path"`
	SearchEnable bool   `yaml:"search_enable"`
}

var config = NewConfig(path.Join(GetRootPath(), "conf/book.yaml"))

func NewConfig(filePath string) *Config {
	log.Printf("configFile: %s", filePath)
	buf, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Fatalf("init config failed: %s", err)
	}

	conf := &Config{}
	err = yaml.Unmarshal(buf, &conf)
	if err != nil {
		log.Fatalf("init config failed: %s", err)
	}

	log.Printf("Config = %#v", conf)

	return conf
}

func GetConfig() *Config {
	return config
}
