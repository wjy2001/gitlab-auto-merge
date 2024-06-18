package conf

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
)

var gconf Config
var rwlock sync.RWMutex

func init() {
	initConfig()
}
func initConfig() {
	//允许初始化没有值
	confBytes, err := os.ReadFile("./conf.json")
	if err != nil {
		var conf Config
		err := UpdateConfig(conf)
		if err != nil {
			log.Fatal(err)
		}
		initConfig()
		return
	}

	err = json.Unmarshal(confBytes, &gconf)
	if err != nil {
		panic(err)
	}
}

type Parameter struct {
	BasicUrl string `json:"basic_url"`
	Token    string `json:"token"`
}
type Config struct {
	Parameter        Parameter      `json:"parameter"`
	ProjectBlacklist map[int]string `json:"project_blacklist"`
}

func UpdateConfig(conf Config) (err error) {
	rwlock.RLock()
	gconf = conf
	rwlock.RUnlock()

	_ = os.Remove("./conf.json")
	by, err := json.Marshal(conf)
	if err != nil {
		err = fmt.Errorf("marshal config error: %w", err)
		return
	}
	err = os.WriteFile("./conf.json", by, 0644)
	if err != nil {
		err = fmt.Errorf("write config error: %w", err)
		return
	}
	return
}

func GetConfig() Config {
	var conf Config
	rwlock.Lock()
	conf = gconf
	rwlock.Unlock()
	return conf
}
