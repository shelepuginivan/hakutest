package i18n

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultI18n(t *testing.T) {
	i := Default()

	i18nBranches := []reflect.Value{
		reflect.ValueOf(i.Server),

		reflect.ValueOf(i.Statistics.App),
		reflect.ValueOf(i.Statistics.Excel),
		reflect.ValueOf(i.Statistics.Image),

		reflect.ValueOf(i.Web.Editor),
		reflect.ValueOf(i.Web.Error),
		reflect.ValueOf(i.Web.Expired),
		reflect.ValueOf(i.Web.Search),
		reflect.ValueOf(i.Web.Submitted),
		reflect.ValueOf(i.Web.Test),
	}

	for _, v := range i18nBranches {
		for i := 0; i < v.NumField(); i++ {
			assert.NotEqual(t, v.Field(i).Interface(), "")
		}
	}
}

func TestNewI18n(t *testing.T) {
	i := New()

	i18nBranches := []reflect.Value{
		reflect.ValueOf(i.Server),

		reflect.ValueOf(i.Statistics.App),
		reflect.ValueOf(i.Statistics.Excel),
		reflect.ValueOf(i.Statistics.Image),

		reflect.ValueOf(i.Web.Editor),
		reflect.ValueOf(i.Web.Error),
		reflect.ValueOf(i.Web.Expired),
		reflect.ValueOf(i.Web.Search),
		reflect.ValueOf(i.Web.Submitted),
		reflect.ValueOf(i.Web.Test),
	}

	for _, v := range i18nBranches {
		for i := 0; i < v.NumField(); i++ {
			assert.NotEqual(t, v.Field(i).Interface(), "")
		}
	}
}
