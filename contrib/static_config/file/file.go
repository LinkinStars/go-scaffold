package file

import (
	"github.com/LinkinStars/go-scaffold/config"
	"github.com/spf13/viper"
)

var _ config.Source = (*file)(nil)

type file struct {
	path string
}

// NewSource new a file source.
func NewSource(path string) config.Source {
	return &file{path: path}
}

func (f *file) Load() error {
	configVip := viper.New()
	configVip.SetConfigFile(f.path)

	configVip.

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}
}

// InitConfig
func InitConfig(path string, c interface{}) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	if err := configVip.Unmarshal(c); err != nil {
		panic(err)
	}
}
