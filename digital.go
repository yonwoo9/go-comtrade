package comtrade

type ComtradeDigital struct {
	Dn   uint32 //数字/状态通道索引号	dn
	ChId string //通道标识		ch_id
	Ph   string //通道相别标识	ph
	Ccbm string //被监视的电路元件	ccbm
	Y    int32  //通道状态	y
}

func (c *ComtradeDigital) GetDn() uint32 {
	return c.Dn
}

func (c *ComtradeDigital) GetChId() string {
	return c.ChId
}

func (c *ComtradeDigital) GetPh() string {
	return c.Ph
}

func (c *ComtradeDigital) GetCcbm() string {
	return c.Ccbm
}

func (c *ComtradeDigital) GetY() int32 {
	return c.Y
}
