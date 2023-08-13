package comtrade

func init() {
	Add(TypeBinary32, func() Parser {
		return &Binary32Data{}
	})
}

// Binary32Data TODO
type Binary32Data struct {
}

func (b *Binary32Data) Parse(filePath string, analogNum, digitalNum, endSamp uint32) (*ComtradeData, error) {
	return &ComtradeData{}, nil
}
