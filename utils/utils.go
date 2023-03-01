package utils

import (
	"errors"
	"github.com/go-ini/ini"
	"log"
	"path/filepath"
)

func ConfigurationLoading(types string, list []string) ([]string, error) {
	absPath, err := filepath.Abs("config.ini")
	if err != nil {
		log.Println(err)
	}

	cfg, err := ini.Load(absPath)
	if err != nil {
		log.Panicln("utils:加载配置文件失败:", err)
	}
	returnList := make([]string, len(list))
	if len(list) != 0 {
		for i := 0; i < len(list); i++ {
			returnList[i] = cfg.Section(types).Key(list[i]).String()
		}
		return returnList, nil
	}

	return returnList, errors.New("utils:配置文件读取失败")
}
