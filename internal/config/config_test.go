package config

import (
	"path"
	"reflect"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestGetConfigDir(t *testing.T) {
	dir := getConfigDir()

	if !strings.HasSuffix(dir, "/hakutest") {
		t.Fail()
	}
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
			if v.Field(i).Interface() == "" {
				t.Fail()
			}
		}
	}
}

func TestPrint(t *testing.T) {
	Init()

	if Print() != nil {
		t.Fail()
	}
}

func TestPrintField(t *testing.T) {
	Init()

	viper.SetConfigFile(path.Join(getConfigDir(), "config.yaml"))
	viper.ReadInConfig()

	keys := viper.AllKeys()

	for _, k := range keys {
		if PrintField(k) != nil {
			t.Fail()
		}
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
			if SetField(k, v.(string)) != nil {
				t.Fail()
			}
		}
	}
}

func TestNegativeSetField(t *testing.T) {
	Init()

	for _, k := range []string{"general", "server", "stats", "stats.excel", "stats.image", "ui", "ui.error", "ui.test"} {
		if SetField(k, "") == nil {
			t.Fail()
		}
	}
}
