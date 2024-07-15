package comtrade

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

func ParseComtradeCfg(filePath string) (*ComtradeCfg, error) {
	cfgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer cfgFile.Close()

	comtradeCfg := &ComtradeCfg{}
	var tempList [][]byte
	content, err := io.ReadAll(cfgFile)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(content, []byte("\n"))

	// read first line
	tempList = bytes.Split(lines[0], []byte(","))
	if len(tempList) < 3 {
		return nil, ErrReadFirstLine
	}
	// station_name, rec_dev_id, rev_year
	comtradeCfg.StationName = ByteToString(tempList[0])
	comtradeCfg.RecDevID = ByteToString(tempList[1])
	if value, err := strconv.ParseUint(ByteToString(tempList[2]), 10, 16); err != nil {
		return nil, err
	} else {
		comtradeCfg.RevYear = uint16(value)
	}

	// read second line
	tempList = bytes.Split(lines[1], []byte(","))
	if len(tempList) < 3 {
		return nil, ErrReadSecondLine
	}

	// total channel number
	if value, err := strconv.ParseUint(ByteToString(tempList[0]), 10, 16); err != nil {
		return nil, err
	} else {
		comtradeCfg.Total = uint32(value)
	}
	if !bytes.Contains(tempList[1], []byte("A")) || !bytes.Contains(tempList[2], []byte("D")) {
		return nil, ErrReadADChannel
	}
	// analog channel total number
	if value, err := strconv.ParseUint(string(bytes.TrimSuffix(bytes.TrimSpace(tempList[1]), []byte("A"))), 10, 64); err != nil {
		return nil, err
	} else {
		comtradeCfg.AnalogNum = uint32(value)
	}
	// digit channel total number
	if value, err := strconv.ParseUint(string(bytes.TrimSuffix(bytes.TrimSpace(tempList[2]), []byte("D"))), 10, 64); err != nil {
		return nil, err
	} else {
		comtradeCfg.DigitalNum = uint32(value)
	}

	// initialize analog and digital channels
	comtradeCfg.Analog = make([]AnalogChan, comtradeCfg.AnalogNum)
	comtradeCfg.Digital = make([]DigitalChan, comtradeCfg.DigitalNum)

	// 读取模拟通道 read analog channels
	for i := 0; i < int(comtradeCfg.AnalogNum); i++ {
		tempList = bytes.Split(lines[2+i], []byte(","))
		if len(tempList) < 10 {
			return nil, ErrReadAnalogChannel
		}
		// 通道索引号 An
		if num, err := strconv.ParseInt(ByteToString(tempList[0]), 10, 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].An = uint32(num)
		}
		// 通道标识 ch_id，通道相标识 ph
		comtradeCfg.Analog[i].ChId = ByteToString(tempList[1])
		comtradeCfg.Analog[i].Ph = ByteToString(tempList[2])
		// 被监视的电路元件 ccbm， 通道单位 uu
		comtradeCfg.Analog[i].Ccbm = ByteToString(tempList[3])
		comtradeCfg.Analog[i].Uu = ByteToString(tempList[4])
		// 通道增益系数 a
		if num, err := strconv.ParseFloat(ByteToString(tempList[5]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].A = float32(num)
		}
		// 通道偏移量 b
		if num, err := strconv.ParseFloat(ByteToString(tempList[6]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].B = float32(num)
		}
		// 从采样时刻开始的通道时滞 skew
		if num, err := strconv.ParseFloat(ByteToString(tempList[7]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].Skew = float32(num)
		}
		// Min Value at current channel
		if num, err := strconv.ParseFloat(ByteToString(tempList[8]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].Min = float32(num)
		}
		// Max Value at current channel
		if num, err := strconv.ParseFloat(ByteToString(tempList[9]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Analog[i].Max = float32(num)
		}

		if len(tempList) > 10 {
			if num, err := strconv.ParseFloat(ByteToString(tempList[10]), 64); err == nil {
				comtradeCfg.Analog[i].Primary = float32(num)
			}
		}
		if len(tempList) > 11 {
			if num, err := strconv.ParseFloat(ByteToString(tempList[11]), 64); err == nil {
				comtradeCfg.Analog[i].Secondary = float32(num)
			}
		}
	}

	// read digit channels
	for i := 0; i < int(comtradeCfg.DigitalNum); i++ {
		tempList = bytes.Split(lines[2+int(comtradeCfg.AnalogNum)+i], []byte(","))
		if len(tempList) < 3 {
			return nil, ErrReadDigitalChannel
		}
		if num, err := strconv.Atoi(ByteToString(tempList[0])); err != nil {
			return nil, err
		} else {
			comtradeCfg.Digital[i].Dn = uint32(num)
		}
		comtradeCfg.Digital[i].ChId = ByteToString(tempList[1])
		comtradeCfg.Digital[i].Ph = ByteToString(tempList[2])

		// checking vector length to avoid IndexError
		if len(tempList) > 3 {
			// channel element (usually null)
			comtradeCfg.Digital[i].Ccbm = ByteToString(tempList[3])
		}
		if len(tempList) > 4 {
			if num, err := strconv.ParseUint(ByteToString(tempList[4]), 10, 64); err != nil {
				return nil, err
			} else {
				comtradeCfg.Digital[i].Y = uint8(num)
			}
		}
	}

	// read frequency
	var line uint32 = 2
	tempList = bytes.Split(lines[line+comtradeCfg.Total], []byte(","))
	if num, err := strconv.ParseFloat(ByteToString(tempList[0]), 64); err != nil {
		return nil, err
	} else {
		comtradeCfg.Lf = float32(num)
		line += 1
	}

	// read sampling rate num
	tempList = bytes.Split(lines[line+comtradeCfg.Total], []byte(","))
	if num, err := strconv.ParseUint(ByteToString(tempList[0]), 10, 64); err != nil {
		return nil, err
	} else {
		comtradeCfg.Nrates = uint16(num)
		line += 1
	}

	// read Sample and endSample
	for i := 0; i < int(comtradeCfg.Nrates); i++ {
		tempList = bytes.Split(lines[line+uint32(i)+comtradeCfg.Total], []byte(","))
		if len(tempList) < 2 {
			return nil, ErrReadSample
		}
		if num, err := strconv.ParseFloat(ByteToString(tempList[0]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.Samp = append(comtradeCfg.Samp, float32(num))
		}
		if num, err := strconv.ParseFloat(ByteToString(tempList[1]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.EndSamp = append(comtradeCfg.EndSamp, uint32(num))
		}
	}
	line += uint32(comtradeCfg.Nrates)

	// read first data time (dd/mm/yyyy,hh:mm:ss.ssssss)
	if start, err := time.Parse(DateTimeLayout, ByteToString(lines[line+comtradeCfg.Total])); err != nil {
		if strings.Contains(err.Error(), MonthOutOfRange) {
			// try to parse date reverse month and day: mm/dd/yyyy
			if start, err = time.Parse(DateTimeLayout2, ByteToString(lines[line+comtradeCfg.Total])); err != nil {
				return nil, err
			} else {
				comtradeCfg.FirstDataTime = start
				line += 1
			}
		}
		if err != nil {
			return nil, err
		}
	} else {
		comtradeCfg.FirstDataTime = start
		line += 1
	}

	// read trigger time (dd/mm/yyyy,hh:mm:ss.ssssss)
	tempList = bytes.Split(lines[line+comtradeCfg.Total], []byte(","))
	if trigger, err := time.Parse(DateTimeLayout, ByteToString(bytes.Join(tempList, []byte(",")))); err != nil {
		if strings.Contains(err.Error(), MonthOutOfRange) {
			// try to parse date reverse month and day
			if trigger, err = time.Parse(DateTimeLayout2, ByteToString(bytes.Join(tempList, []byte(",")))); err != nil {
				return nil, err
			} else {
				comtradeCfg.TriggerTime = trigger
				line += 1
			}
		}
		if err != nil {
			return nil, err
		}
	} else {
		comtradeCfg.TriggerTime = trigger
		line += 1
	}

	// read data file type
	tempList = bytes.Split(lines[line+comtradeCfg.Total], []byte(","))
	comtradeCfg.Ft = ByteToString(tempList[0])

	// read time multiplication factor
	line += 1
	tempList = bytes.Split(lines[line+comtradeCfg.Total], []byte(","))
	if !bytes.Equal(tempList[0], []byte("")) {
		if num, err := strconv.ParseFloat(ByteToString(tempList[0]), 64); err != nil {
			return nil, err
		} else {
			comtradeCfg.TimeMult = float32(num)
		}
	} else {
		comtradeCfg.TimeMult = 1
		line += 1
	}
	return comtradeCfg, nil
}

func ByteToString(b []byte) string {
	return strings.TrimSpace(string(b))
}
