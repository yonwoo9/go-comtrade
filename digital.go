package comtrade

type ComtradeDigitalChan struct {
	Dn   uint32 //数字/状态通道索引号	dn
	ChId string //通道标识		ch_id
	Ph   string //通道相别标识	ph
	Ccbm string //被监视的电路元件	ccbm
	Y    uint8  //通道状态	y
}

func (c *ComtradeDigitalChan) GetDn() uint32 {
	return c.Dn
}

func (c *ComtradeDigitalChan) GetChId() string {
	return c.ChId
}

func (c *ComtradeDigitalChan) GetPh() string {
	return c.Ph
}

func (c *ComtradeDigitalChan) GetCcbm() string {
	return c.Ccbm
}

func (c *ComtradeDigitalChan) GetY() uint8 {
	return c.Y
}
