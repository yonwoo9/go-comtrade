package comtrade

import (
	"reflect"
	"testing"
)

func TestBinaryData_Parse(t *testing.T) {
	filePath := "testdata/test1.cfg"
	cfg, err := ParseComtradeCfg(filePath)
	if err != nil {
		t.Errorf("ParseComtradeCfg() error = %v", err)
		return
	}

	b := &BinaryData{}
	dat, err := b.Parse("testdata/test1.dat", cfg.AnalogNum, cfg.DigitalNum, cfg.EndSamp[len(cfg.EndSamp)-1])
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}
	if reflect.TypeOf(dat) != reflect.TypeOf(&Data{}) {
		t.Errorf("Parse() got = %v", dat)
	}
	t.Log(dat)
}
