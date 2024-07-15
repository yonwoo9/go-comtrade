package comtrade

import (
	"time"
)

type ComtradeCfg struct {
	StationName   string        `json:"station_name"`    //厂站名称
	RecDevID      string        `json:"rec_dev_id"`      //记录设备ID
	RevYear       uint16        `json:"rev_year"`        //COMTRADE版本年份	1991、1999、2013
	Total         uint32        `json:"total"`           //总通道数
	AnalogNum     uint32        `json:"analog_num"`      //模拟通道数
	DigitalNum    uint32        `json:"digital_num"`     //数字/状态通道数
	Analog        []AnalogChan  `json:"analog"`          //模拟通道
	Digital       []DigitalChan `json:"digital"`         //数字/状态通道
	Lf            float32       `json:"lf"`              //标称频率
	Nrates        uint16        `json:"nrates"`          //采样率个数
	Samp          []float32     `json:"samp"`            //采样率
	EndSamp       []uint32      `json:"end_samp"`        //最末采样序号
	FirstDataTime time.Time     `json:"first_data_time"` //第一条数据时间
	TriggerTime   time.Time     `json:"trigger_time"`    //采样触发时间
	Ft            string        `json:"ft"`              //数据文件类型，ASCII、BINARY、BINARY32、FLOAT32

	// 2017
	TimeMult  float32 `json:"time_mult"`  //时间倍率因子
	TimeCode  string  `json:"time_code"`  //时间编码
	LocalCode string  `json:"local_code"` //本地编码
	TmqCode   uint8   `json:"tmq_code"`   //采样时间品质
	Leapsec   uint8   `json:"leapsec"`    //闰秒标识符
}

func (cc *ComtradeCfg) GetStationName() string {
	return cc.StationName
}

func (cc *ComtradeCfg) GetRecDevID() string {
	return cc.RecDevID
}

func (cc *ComtradeCfg) GetRevYear() uint16 {
	return cc.RevYear
}

func (cc *ComtradeCfg) GetTotal() uint32 {
	return cc.Total
}

func (cc *ComtradeCfg) GetAnalogNum() uint32 {
	return cc.AnalogNum
}

func (cc *ComtradeCfg) GetDigitalNum() uint32 {
	return cc.DigitalNum
}

func (cc *ComtradeCfg) GetLf() float32 {
	return cc.Lf
}

func (cc *ComtradeCfg) GetNrates() uint16 {
	return cc.Nrates
}

func (cc *ComtradeCfg) GetSamp() []float32 {
	return cc.Samp
}

func (cc *ComtradeCfg) GetEndSamp() []uint32 {
	return cc.EndSamp
}

func (cc *ComtradeCfg) GetTriggerTime() time.Time {
	return cc.TriggerTime
}

func (cc *ComtradeCfg) GetFirstDataTime() time.Time {
	return cc.FirstDataTime
}

func (cc *ComtradeCfg) GetFt() string {
	return cc.Ft
}

func (cc *ComtradeCfg) GetTimeMult() float32 {
	return cc.TimeMult
}

func (cc *ComtradeCfg) GetTimeCode() string {
	return cc.TimeCode
}

func (cc *ComtradeCfg) GetLocalCode() string {
	return cc.LocalCode
}

func (cc *ComtradeCfg) GetTmqCode() uint8 {
	return cc.TmqCode
}

func (cc *ComtradeCfg) GetLeapsec() uint8 {
	return cc.Leapsec
}

func (cc *ComtradeCfg) GetAnalog() []AnalogChan {
	return cc.Analog
}

func (cc *ComtradeCfg) GetDigital() []DigitalChan {
	return cc.Digital
}
