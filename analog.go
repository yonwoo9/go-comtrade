package comtrade

type ComtradeAnalog struct {
	An        string  //模拟通道索引号 an
	ChId      string  //通道标识		 ch_id
	Ph        string  //通道相别标识	ph
	Ccbm      string  //被监视的电路元件	ccbm
	Uu        string  //通道单位	uu
	A         float64 //通道增益系数	a
	B         float64 //通道偏移量	b
	Skew      float64 //通道时滞	skew
	Min       float64 //通道最小值	min
	Max       float64 //通道最大值	max
	Primary   float64 //一次系数	primary
	Secondary float64 //二次系数	secondary
	PS        string  //一次二次标识	ps
}

func (c *ComtradeAnalog) GetAn() string {
	return c.An
}

func (c *ComtradeAnalog) GetChId() string {
	return c.ChId
}

func (c *ComtradeAnalog) GetPh() string {
	return c.Ph
}

func (c *ComtradeAnalog) GetCcbm() string {
	return c.Ccbm
}

func (c *ComtradeAnalog) GetUu() string {
	return c.Uu
}

func (c *ComtradeAnalog) GetA() float64 {
	return c.A
}

func (c *ComtradeAnalog) GetB() float64 {
	return c.B
}

func (c *ComtradeAnalog) GetSkew() float64 {
	return c.Skew
}

func (c *ComtradeAnalog) GetMin() float64 {
	return c.Min
}

func (c *ComtradeAnalog) GetMax() float64 {
	return c.Max
}

func (c *ComtradeAnalog) GetPrimary() float64 {
	return c.Primary
}

func (c *ComtradeAnalog) GetSecondary() float64 {
	return c.Secondary
}

func (c *ComtradeAnalog) GetPS() string {
	return c.PS
}
