package comtrade

type Comtrade struct {
	Conf *ComtradeCfg
	Data *Data
}

func ParseComtrade(configurePath, datafilePath string) (comtrade *Comtrade, err error) {
	c := &Comtrade{}
	c.Conf, err = c.parseComtradeConfig(configurePath)
	if err != nil {
		return nil, err
	}
	c.Data, err = c.parseComtradeData(datafilePath)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Comtrade) parseComtradeConfig(configurePath string) (*ComtradeCfg, error) {
	comtradeConfig, err := ParseComtradeCfg(configurePath)
	if err != nil {
		return nil, err
	}
	return comtradeConfig, nil
}

func (c *Comtrade) parseComtradeData(datafilePath string) (*Data, error) {
	creator, ok := Parsers[c.Conf.Ft]
	if !ok {
		return nil, ErrUnknownTypeOfData
	}
	parser := creator()
	comtradeData, err := parser.Parse(datafilePath, c.Conf.GetAnalogNum(), c.Conf.GetDigitalNum(), c.Conf.GetEndSamp()[len(c.Conf.GetEndSamp())-1])
	if err != nil {
		return nil, err
	}
	return comtradeData, nil
}
