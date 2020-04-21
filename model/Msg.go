package model

type Msg struct {
	MsgId   int64
	FromId  int64
	Cmd     int
	GroupId int64
	DstId   int64
}
