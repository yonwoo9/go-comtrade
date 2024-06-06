package comtrade

import (
	"reflect"
	"testing"
)

func TestAsciiData_Parse(t *testing.T) {
	filePath := "testdata/test2.cfg"
	cfg, err := ParseComtradeCfg(filePath)
	if err != nil {
		t.Errorf("ParseComtradeCfg() error = %v", err)
		return
	}

	a := &AsciiData{}
	dat, err := a.Parse("testdata/test2.dat", cfg.AnalogNum, cfg.DigitalNum, cfg.EndSamp[len(cfg.EndSamp)-1])
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}
	if reflect.TypeOf(dat) != reflect.TypeOf(&Data{}) {
		t.Errorf("Parse() got = %v", dat)
	}
	t.Log(dat)
}
