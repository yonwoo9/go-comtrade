package comtrade

import (
	"encoding/binary"
	"math"
	"os"
)

// BinaryData is the binary data of a comtrade data file

func init() {
	Add(TypeBinary, func() Parser {
		return &BinaryData{}
	})
}

//读取数据文件
//1、采样序号和时标，均以四字节，无符号二进制格式存储
//2、模拟通道采样数据，binary两个字节，binary32四个字节二进制补码形式存储
//3、每16个状态通道以两个字节一组存储，字的最低为对应该组16个状态通道中最小编号通道
//4、每次采样字节数=（模拟通道数 * 每个采样数据占据字节数）+（状态通道数 / 16 * 2） + 4 + 4
//5、每个采样数据占据字节数=2或4，取决于数据格式，2表示binary，4表示binary32

//05，667，-760，1274，72，61，-140，-502，0，0，0，0，1，1
//05 00 00 00
//9B 02 00 00	667
//08 FD 	   -760
//FA 04 	    1274
//48 00			72
//3D 00 		61
//74 FF			-140
//0A FE 	    -502
//30 00			0，0，0，0，1，1

type BinaryData struct {
}

func (b *BinaryData) Parse(filePath string, analogNum, digitalNum, endSamp uint32) (*Data, error) {
	comtradeData := &Data{}

	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	comtradeData.AnalogData = make([]AnalogData, int(endSamp))
	comtradeData.DigitalData = make([]DigitalData, int(endSamp))

	for i := 0; i < int(endSamp); i++ {
		comtradeData.AnalogData[i].Data = make([]int32, analogNum)
		comtradeData.DigitalData[i].Data = make([]uint8, digitalNum)
	}

	// 计算状态通道组数
	digitalCount := int(math.Ceil(float64(digitalNum) / 16.0))
	remainder := int(digitalNum % 16)

	for i := 0; i < int(endSamp); i++ {
		var (
			sampleIndex uint32
			timestamp   uint32
		)
		// 解析采样序号
		if err := binary.Read(file, binary.LittleEndian, &sampleIndex); err != nil {
			return nil, err
		}
		comtradeData.AnalogData[i].N = sampleIndex
		comtradeData.DigitalData[i].N = sampleIndex

		// 解析采样时标
		if err = binary.Read(file, binary.LittleEndian, &timestamp); err != nil {
			return nil, err
		}
		comtradeData.AnalogData[i].Timestamp = timestamp
		comtradeData.DigitalData[i].Timestamp = timestamp

		// 解析模拟通道采样数据
		for m := 0; m < int(analogNum); m++ {
			var tmp int16
			if err := binary.Read(file, binary.LittleEndian, &tmp); err != nil {
				return nil, err
			} else {
				comtradeData.AnalogData[i].Data[m] = int32(tmp)
			}
		}

		stateData := make([]uint16, digitalCount)
		for n, datum := range stateData {
			data := datum
			if err := binary.Read(file, binary.LittleEndian, &data); err != nil {
				return nil, err
			}
			stateData[n] = data
		}
		for h, datum := range stateData {
			if remainder != 0 && h == digitalCount-1 {
				for j := 0; j < remainder; j++ {
					comtradeData.DigitalData[i].Data[(h*16)+j] = uint8(uint((datum >> uint(j)) & 0x0001))
				}
				break
			}
			for k := 0; k < 16; k++ {
				comtradeData.DigitalData[i].Data[(h*16)+k] = uint8(uint((datum >> uint(k)) & 0x0001))
			}
		}
	}
	return comtradeData, nil
}
