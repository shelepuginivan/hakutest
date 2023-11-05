package config

import (
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetConfigDir(t *testing.T) {
	assert.True(t, strings.HasSuffix(getConfigDir(), "/hakutest"))
}

func TestInitConfig(t *testing.T) {
	c := Init()

	var (
		g     = reflect.ValueOf(c.General)
		s     = reflect.ValueOf(c.Server)
		s_e   = reflect.ValueOf(c.Statistics.Excel)
		s_i   = reflect.ValueOf(c.Statistics.Image)
		ui_ed = reflect.ValueOf(c.Ui.Editor)
		ui_er = reflect.ValueOf(c.Ui.Error)
		ui_t  = reflect.ValueOf(c.Ui.Test)
	)

	for _, v := range []reflect.Value{g, s, s_e, s_i, ui_ed, ui_er, ui_t} {
		for i := 0; i < v.NumField(); i++ {
			assert.NotEqual(t, v.Field(i).Interface(), "")
		}
	}
}

func TestPrint(t *testing.T) {
	Init()

	assert.Nil(t, Print())
}

func TestPrintField(t *testing.T) {
	Init()

	viper.SetConfigFile(path.Join(getConfigDir(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		assert.Nil(t, PrintField(k))
	}
}

func TestSetField(t *testing.T) {
	Init()

	viper.SetConfigFile(path.Join(getConfigDir(), "config.yaml"))
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
	Init()

	for _, k := range []string{"general", "server", "stats", "stats.excel", "stats.image", "ui", "ui.error", "ui.test"} {
		assert.Error(t, SetField(k, ""))
	}
}