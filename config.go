package comtrade

import "time"

type ComtradeCfg struct {
	StationName   string    //厂站名称	station_name
	RecDevID      string    //记录设备ID	rec_dev_id
	RevYear       int32     //COMTRADE版本年份	rev_year
	Total         int32     //总通道数	total
	AnalogNum     int32     //模拟通道数	analog_num
	DigtalNum     int32     //数字/状态通道数	digital_num
	Lf            int32     //标称频率	lf
	Nrates        int32     //采样率个数	nrates
	Samp          int32     //采样率	samp
	EndSamp       int32     //最末采样序号	end_samp
	TriggerTime   time.Time //采样触发时间	trigger_time
	FirstDataTime time.Time //第一条数据时间	first_data_time
	Ft            string    //数据文件类型，ASCII或者BINARY或者BINARY32,FLOAT32	ft

	// 2017
	TimeMult  float64 //时间倍率因子	time_mult
	TimeCode  int32   //时间编码	time_code
	LocalCode int32   //本地编码	local_code
	TmqCode   int32   //采样时间品质	tmq_code
	Leapsec   int32   //闰秒标识符	leapsec
}

func ParseConfigFile(filePath string) (*ComtradeCfg, error) {
	//TODO
	return nil, nil
}

func (c *ComtradeCfg) GetStationName() string {
	return c.StationName
}

func (c *ComtradeCfg) GetRecDevID() string {
	return c.RecDevID
}

func (c *ComtradeCfg) GetRevYear() int32 {
	return c.RevYear
}

func (c *ComtradeCfg) GetTotal() int32 {
	return c.Total
}

func (c *ComtradeCfg) GetAnalogNum() int32 {
	return c.AnalogNum
}

func (c *ComtradeCfg) GetDigtalNum() int32 {
	return c.DigtalNum
}

func (c *ComtradeCfg) GetLf() int32 {
	return c.Lf
}

func (c *ComtradeCfg) GetNrates() int32 {
	return c.Nrates
}

func (c *ComtradeCfg) GetSamp() int32 {
	return c.Samp
}

func (c *ComtradeCfg) GetEndSamp() int32 {
	return c.EndSamp
}

func (c *ComtradeCfg) GetTriggerTime() time.Time {
	return c.TriggerTime
}

func (c *ComtradeCfg) GetFirstDataTime() time.Time {
	return c.FirstDataTime
}

func (c *ComtradeCfg) GetFt() string {
	return c.Ft
}

func (c *ComtradeCfg) GetTimeMult() float64 {
	return c.TimeMult
}

func (c *ComtradeCfg) GetTimeCode() int32 {
	return c.TimeCode
}

func (c *ComtradeCfg) GetLocalCode() int32 {
	return c.LocalCode
}

func (c *ComtradeCfg) GetTmqCode() int32 {
	return c.TmqCode
}

func (c *ComtradeCfg) GetLeapsec() int32 {
	return c.Leapsec
}
