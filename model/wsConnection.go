/**
*
*@author 吴昊轩
*@create 2020-04-2118:45
 */
package model

import (
	"github.com/gorilla/websocket"
	"time"
)

type WsConnection struct {
	Uid               int64
	wsSocket          *websocket.Conn
	recvChan          chan *Msg
	sendChan          chan *Msg
	lastHeartbeatTime time.Time
	groupIds          map[string]bool
}
