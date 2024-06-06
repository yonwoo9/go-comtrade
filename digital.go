package comtrade

// DigitalChan 数字通道
type DigitalChan struct {
	Dn   uint32 `json:"dn"`    //数字/状态通道索引号
	ChId string `json:"ch_id"` //通道标识
	Ph   string `json:"ph"`    //通道相别标识
	Ccbm string `json:"ccbm"`  //被监视的电路元件
	Y    uint8  `json:"y"`     //通道状态
}

func (c *DigitalChan) GetDn() uint32 {
	return c.Dn
}

func (c *DigitalChan) GetChId() string {
	return c.ChId
}

func (c *DigitalChan) GetPh() string {
	return c.Ph
}

func (c *DigitalChan) GetCcbm() string {
	return c.Ccbm
}

func (c *DigitalChan) GetY() uint8 {
	return c.Y
}
