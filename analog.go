package comtrade

// AnalogChan 模拟通道
type AnalogChan struct {
	An        uint32  `json:"an"`        //模拟通道索引号
	ChId      string  `json:"ch_id"`     //通道标识
	Ph        string  `json:"ph"`        //通道相别标识
	Ccbm      string  `json:"ccbm"`      //被监视的电路元件
	Uu        string  `json:"uu"`        //通道单位
	A         float32 `json:"a"`         //通道增益系数
	B         float32 `json:"b"`         //通道偏移量
	Skew      float32 `json:"skew"`      //通道时滞
	Min       float32 `json:"min"`       //通道最小值
	Max       float32 `json:"max"`       //通道最大值
	Primary   float32 `json:"primary"`   //一次系数
	Secondary float32 `json:"secondary"` //二次系数
	PS        string  `json:"ps"`        //一次二次标识
}

func (c *AnalogChan) GetAn() uint32 {
	return c.An
}

func (c *AnalogChan) GetChId() string {
	return c.ChId
}

func (c *AnalogChan) GetPh() string {
	return c.Ph
}

func (c *AnalogChan) GetCcbm() string {
	return c.Ccbm
}

func (c *AnalogChan) GetUu() string {
	return c.Uu
}

func (c *AnalogChan) GetA() float32 {
	return c.A
}

func (c *AnalogChan) GetB() float32 {
	return c.B
}

func (c *AnalogChan) GetSkew() float32 {
	return c.Skew
}

func (c *AnalogChan) GetMin() float32 {
	return c.Min
}

func (c *AnalogChan) GetMax() float32 {
	return c.Max
}

func (c *AnalogChan) GetPrimary() float32 {
	return c.Primary
}

func (c *AnalogChan) GetSecondary() float32 {
	return c.Secondary
}

func (c *AnalogChan) GetPS() string {
	return c.PS
}
