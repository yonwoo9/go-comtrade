package comtrade

type Comtrade struct {
	Conf *ComtradeCfg
	Data *ComtradeData
}

func ParseComtrade(configfilePath, datafilePath string) (comtrade *Comtrade, err error) {
	c := &Comtrade{}
	c.Conf, err = c.parseComtradeConfig(configfilePath)
	if err != nil {
		return nil, err
	}
	c.Data, err = c.parseComtradeData(datafilePath)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Comtrade) parseComtradeConfig(configfilePath string) (*ComtradeCfg, error) {
	comtradeConfig, err := ParseComtradeCfg(configfilePath)
	if err != nil {
		return nil, err
	}
	return comtradeConfig, nil
}

func (c *Comtrade) parseComtradeData(datafilePath string) (*ComtradeData, error) {
	creator, ok := Parsers[c.Conf.Ft]
	if !ok {
		return nil, ErrUnknowTypeOfData
	}
	parser := creator()
	comtradeData, err := parser.Parse(datafilePath, c.Conf.GetAnalogNum(), c.Conf.GetDigtalNum(), c.Conf.GetEndSamp()[len(c.Conf.GetEndSamp())-1])
	if err != nil {
		return nil, err
	}
	return comtradeData, nil
}
