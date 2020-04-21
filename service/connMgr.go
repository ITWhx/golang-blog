/**
*
*@author 吴昊轩
*@create 2020-04-2119:09
 */
package service

import (
	"golang-blog/model"
	"sync"
)

type ConnMgr struct {
	buckets []*bucket
}

type bucket struct {
	RwMutex       sync.RWMutex
	idx           int
	wsConnections map[int64]*model.WsConnection
}

func (b *bucket) addWsConnection(ws *model.WsConnection) {
	b.RwMutex.RLock()
	defer b.RwMutex.Unlock()
	b.wsConnections[ws.Uid] = ws
}
