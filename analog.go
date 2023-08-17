package comtrade

type ComtradeAnalogChan struct {
	An        uint32  //模拟通道索引号 an
	ChId      string  //通道标识		 ch_id
	Ph        string  //通道相别标识	ph
	Ccbm      string  //被监视的电路元件	ccbm
	Uu        string  //通道单位	uu
	A         float32 //通道增益系数	a
	B         float32 //通道偏移量	b
	Skew      float32 //通道时滞	skew
	Min       float32 //通道最小值	min
	Max       float32 //通道最大值	max
	Primary   float32 //一次系数	primary
	Secondary float32 //二次系数	secondary
	PS        string  //一次二次标识	ps
}

func (c *ComtradeAnalogChan) GetAn() uint32 {
	return c.An
}

func (c *ComtradeAnalogChan) GetChId() string {
	return c.ChId
}

func (c *ComtradeAnalogChan) GetPh() string {
	return c.Ph
}

func (c *ComtradeAnalogChan) GetCcbm() string {
	return c.Ccbm
}

func (c *ComtradeAnalogChan) GetUu() string {
	return c.Uu
}

func (c *ComtradeAnalogChan) GetA() float32 {
	return c.A
}

func (c *ComtradeAnalogChan) GetB() float32 {
	return c.B
}

func (c *ComtradeAnalogChan) GetSkew() float32 {
	return c.Skew
}

func (c *ComtradeAnalogChan) GetMin() float32 {
	return c.Min
}

func (c *ComtradeAnalogChan) GetMax() float32 {
	return c.Max
}

func (c *ComtradeAnalogChan) GetPrimary() float32 {
	return c.Primary
}

func (c *ComtradeAnalogChan) GetSecondary() float32 {
	return c.Secondary
}

func (c *ComtradeAnalogChan) GetPS() string {
	return c.PS
}
