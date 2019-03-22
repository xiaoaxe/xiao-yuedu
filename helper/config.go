// 读取配置文件
// author: baoqiang
// time: 2019/3/21 下午8:05
package helper

import (
	"io/ioutil"
	"github.com/go-yaml/yaml"
	"log"
	"path"
)

type Config struct {
	Sqlite3 string `yaml:"sqlite3"`
	Wukong  Wukong `yaml:"wukong"`
}

type Wukong struct {
	DictFile  string `yaml:"dict_file"`
	IndexPath string `yaml:"index_path"`
}

var config = NewConfig(path.Join(GetRootPath(), "conf/book.yaml"))

func NewConfig(filepath string) *Config {
	log.Printf("configFile: %s", filepath)
	buf, err := ioutil.ReadFile(filepath)
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
