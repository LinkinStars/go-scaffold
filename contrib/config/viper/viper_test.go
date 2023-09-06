package viper

import (
	"fmt"
	"testing"
)

type AllConf struct {
	Base struct {
		WebPort string
	}
	Logger struct {
		Level string
	}
}

func TestNewStaticConfigParser(t *testing.T) {
	configParser := NewStaticConfigParser("./testdata/conf.yaml")
	c := &AllConf{}
	err := configParser.LoadAndSet(c)
	if err != nil {
		t.Error(err)
	}
	if c.Base.WebPort != "8080" || c.Logger.Level != "debug" {
		t.Error(fmt.Errorf("load config failed, config is %+v", c))
	}
}
