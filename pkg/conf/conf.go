package conf

import (
	"expert-back/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DownloadFileInfo struct {
	Path string
	Name string
}

type SystemConf struct {
	System struct {
		Listen string
	}
	File struct {
		Download struct {
			Recommend struct {
				DownloadFileInfo	`yaml:"info"`
			}
		}
	}
	Database struct {
		Name string
		Connection string
	}
}

func Init(path string) {
	getDefault()
	file, err := ioutil.ReadFile(path)
	if err != nil {
		util.Log().Panic("无法读取配置文件 '%s': %s", path, err)
	}
	err = yaml.Unmarshal(file, &SystemConfig)
	if err != nil {
		util.Log().Panic("解析配置文件失败: %s", err)
	}
	// util.Log().Info("配置文件内容: %v", SystemConfig)
}

