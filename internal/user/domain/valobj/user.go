package valobj

type TraceType int8

const (
	TraceTypeReg TraceType = iota
	TraceTypeLogin
	TraceTypeOut
	TraceTypeEdit
	TraceTypeDel
)
