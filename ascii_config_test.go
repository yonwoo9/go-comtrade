package comtrade

import (
	"reflect"
	"testing"
)

func TestParseComtradeCfg(t *testing.T) {
	filePath := "testdata/test1.cfg"
	got, err := ParseComtradeCfg(filePath)
	if err != nil {
		t.Errorf("ParseComtradeCfg() error = %v", err)
		return
	}
	if reflect.TypeOf(got) != reflect.TypeOf(&ComtradeCfg{}) {
		t.Errorf("ParseComtradeCfg() got = %v", got)
	}
	t.Log(got.StationName)
}
