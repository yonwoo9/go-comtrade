package comtrade

import (
	"bytes"
	"io"
	"os"
	"strconv"
)

// Parse AsciiData AsciiData is the ascii data of a comtrade data file.

func init() {
	Add(TypeAscii, func() Parser {
		return &AsciiData{}
	})
}

type AsciiData struct {
}

func (a *AsciiData) Parse(filePath string, analogNum, digitalNum, endSamp uint32) (*ComtradeData, error) {
	comtradeData := &ComtradeData{}
	datFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer datFile.Close()

	content, err := io.ReadAll(datFile)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(content, []byte("\n"))
	if len(lines)-1 != int(endSamp) {
		return nil, ErrInvalidFile
	}

	comtradeData.AnalogData = make([]ComtradeAnalogData, int(endSamp))
	comtradeData.DigitalData = make([]ComtradeDigitalData, int(endSamp))

	for i := 0; i < int(endSamp); i++ {
		comtradeData.AnalogData[i].Data = make([]int32, analogNum)
		comtradeData.DigitalData[i].Data = make([]uint8, digitalNum)
	}

	total := 1 + 1 + int(analogNum) + int(digitalNum)
	var tempList [][]byte
	for i := 0; i < int(endSamp); i++ {
		tempList = bytes.Split(lines[i], []byte(","))
		if len(tempList) != total {
			return nil, ErrReadFirstLine
		}

		if value, err := strconv.Atoi(ByteToString(tempList[0])); err != nil {
			return nil, err
		} else {
			comtradeData.AnalogData[i].N = uint32(value)
			comtradeData.DigitalData[i].N = uint32(value)
		}

		if value, err := strconv.Atoi(ByteToString(tempList[1])); err != nil {
			return nil, err
		} else {
			comtradeData.AnalogData[i].Timestamp = uint32(value)
			comtradeData.DigitalData[i].Timestamp = uint32(value)
		}

		for j := 0; j < int(analogNum); j++ {
			if value, err := strconv.ParseInt(ByteToString(tempList[j+2]), 10, 64); err != nil {
				return nil, err
			} else {
				comtradeData.AnalogData[i].Data[j] = int32(value)
			}
		}

		for j := 0; j < int(digitalNum); j++ {
			if value, err := strconv.ParseInt(ByteToString(tempList[j+2+int(analogNum)]), 10, 64); err != nil {
				return nil, err
			} else {
				comtradeData.DigitalData[i].Data[j] = uint8(value)
			}
		}
	}

	return comtradeData, nil
}
