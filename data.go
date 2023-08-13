package comtrade

// ComtradeData is the data of a comtrade file.
type ComtradeData struct {
	AnalogData  []ComtradeAnalogData
	DigitalData []ComtradeDigitalData
}

// ComtradeAnalogData is the analog data of a comtrade file.
type ComtradeAnalogData struct {
	N         uint32  //模拟通道序号 n
	Timestamp uint32  //模拟通道时标 timestamp
	Data      []int32 //模拟通道数据 data
}

// ComtradeDigitalData is the digital data of a comtrade file.
type ComtradeDigitalData struct {
	N         uint32  //数字通道序号 n
	Timestamp uint32  //数字通道时标 timestamp
	Data      []uint8 //数字通道数据 data
}

func (cd *ComtradeData) GetAnalogData() []ComtradeAnalogData {
	return cd.AnalogData
}

func (cd *ComtradeData) GetDigitalData() []ComtradeDigitalData {
	return cd.DigitalData
}
