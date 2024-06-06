package comtrade

func init() {
	Add(TypeFloat32, func() Parser {
		return &Float32Data{}
	})
}

// Float32Data TODO
type Float32Data struct {
}

func (f *Float32Data) Parse(filePath string, analogNum, digitalNum, endSamp uint32) (*Data, error) {
	return &Data{}, nil
}
