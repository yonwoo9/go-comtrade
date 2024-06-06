package comtrade

import (
	"time"
)

type ComtradeCfg struct {
	StationName   string        //厂站名称	station_name
	RecDevID      string        //记录设备ID	rec_dev_id
	RevYear       uint16        //COMTRADE版本年份	rev_year	1991、1999、2013
	Total         uint32        //总通道数	total
	AnalogNum     uint32        //模拟通道数	analog_num
	DigitalNum    uint32        //数字/状态通道数	digital_num
	Analog        []AnalogChan  //模拟通道	analog
	Digital       []DigitalChan //数字/状态通道	digital
	Lf            float32       //标称频率	lf
	Nrates        uint16        //采样率个数	nrates
	Samp          []float32     //采样率	samp
	EndSamp       []uint32      //最末采样序号	end_samp
	FirstDataTime time.Time     //第一条数据时间	first_data_time
	TriggerTime   time.Time     //采样触发时间	trigger_time
	Ft            string        //数据文件类型，ASCII、BINARY、BINARY32、FLOAT32	ft

	// 2017
	TimeMult  float32 //时间倍率因子	time_mult
	TimeCode  string  //时间编码	time_code
	LocalCode string  //本地编码	local_code
	TmqCode   uint8   //采样时间品质	tmq_code
	Leapsec   uint8   //闰秒标识符	leapsec
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

func (cc *ComtradeCfg) GetDigtalNum() uint32 {
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
