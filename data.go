package comtrade

// Data is the data of a comtrade file.
type Data struct {
	AnalogData  []AnalogData
	DigitalData []DigitalData
}

// AnalogData is the analog data of a comtrade file.
type AnalogData struct {
	N         uint32  `json:"n"`         //模拟通道序号
	Timestamp uint32  `json:"timestamp"` //模拟通道时标
	Data      []int32 `json:"data"`      //模拟通道数据
}

// DigitalData is the digital data of a comtrade file.
type DigitalData struct {
	N         uint32  `json:"n"`         //数字通道序号
	Timestamp uint32  `json:"timestamp"` //数字通道时标
	Data      []uint8 `json:"data"`      //数字通道数据
}

func (cd *Data) GetAnalogData() []AnalogData {
	return cd.AnalogData
}

func (cd *Data) GetDigitalData() []DigitalData {
	return cd.DigitalData
}
