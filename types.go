package comtrade

const (
	DateTimeLayout  = "02/01/2006,15:04:05.000000" // dd/mm/yyyy,HH:MM:SS.mmmmmm
	DateTimeLayout2 = "01/02/2006,15:04:05.000000" // mm/dd/yyyy,HH:MM:SS.mmmmmm

	Rcv1991 = 1991
	Rcv1999 = 1999
	Rcv2013 = 2013
	Rcv2017 = 2017

	TypeAscii    = "ASCII"
	TypeBinary   = "BINARY"
	TypeBinary32 = "BINARY32"
	TypeFloat32  = "FLOAT32"

	NoLeapSecInRecord                = 0 //记录中没有闰秒 no leap seconds are included in the record
	AddLeapSecToRecord               = 1 //在记录中增加闰秒 add leap seconds to the record
	RemoveLeapSecFromRecord          = 2 //在记录中删除闰秒 remove leap seconds from the record
	ClockSourceNotDefinedLeapSecFunc = 3 //时钟源没有闰秒功能 the clock source does not have leap second functionality.

	MonthOutOfRange = "month out of range"
)

type Parser interface {
	Parse(filePath string, analogNum, digitalNum, endSamp uint32) (*Data, error)
}
