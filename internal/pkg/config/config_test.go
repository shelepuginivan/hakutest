package config

import (
	"path/filepath"
	"reflect"
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/directories"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

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

func TestPrint(t *testing.T) {
	New()

	assert.Nil(t, Print())
}

func TestPrintField(t *testing.T) {
	New()

	viper.SetConfigFile(filepath.Join(directories.Config(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		assert.Nil(t, PrintField(k))
	}
}

func TestSetField(t *testing.T) {
	New()

	viper.SetConfigFile(filepath.Join(directories.Config(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		v := viper.Get(k)

		if _, ok := v.(string); ok {
			assert.Nil(t, SetField(k, v.(string)))
		}
	}
}

func TestNegativeSetField(t *testing.T) {
	New()

	for _, k := range []string{"general", "server", "stats", "stats.excel", "stats.image", "ui", "ui.editor", "ui.error", "ui.expired", "ui.test"} {
		assert.Error(t, SetField(k, ""))
	}
}
