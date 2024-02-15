package config

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	New()
	m.Run()
}

func TestDefaultConfig(t *testing.T) {
	c := Default()

	configBranches := []reflect.Value{
		reflect.ValueOf(c.General),
		reflect.ValueOf(c.Server),
	}

	for _, v := range configBranches {
		for i := 0; i < v.NumField(); i++ {
			assert.NotEqual(t, v.Field(i).Interface(), "")
		}
	}
}

func TestNewConfig(t *testing.T) {
	c := New()

	configBranches := []reflect.Value{
		reflect.ValueOf(c.General),
		reflect.ValueOf(c.Server),
	}

	for _, v := range configBranches {
		for i := 0; i < v.NumField(); i++ {
			assert.NotEqual(t, v.Field(i).Interface(), "")
		}
	}
}

func TestConfigService_PrintConfig(t *testing.T) {
	assert.Nil(t, NewService().PrintConfig())
}

func TestConfigService_PrintField(t *testing.T) {
	c := NewService()
	viper.SetConfigFile(filepath.Join(directories.Config(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		assert.Nil(t, c.PrintField(k))
	}
}

func TestConfigService_GetField(t *testing.T) {
	c := NewService()
	viper.SetConfigFile(filepath.Join(directories.Config(), "config.yaml"))
	viper.ReadInConfig()

	for key, value := range viper.AllSettings() {
		assert.Equal(t, value, c.GetField(key))
	}
}

func TestConfigService_SetField(t *testing.T) {
	c := NewService()
	viper.SetConfigFile(filepath.Join(directories.Config(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		v := viper.Get(k)

		if _, ok := v.(string); ok {
			assert.Nil(t, c.SetField(k, v.(string)))
		}
	}
}
