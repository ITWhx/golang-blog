/**
*
*@author 吴昊轩
*@create 2020-04-2119:37
 */
package handler

import (
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

var (
	G_wsServer *WSServer

	wsUpgrader = websocket.Upgrader{
		// 允许所有CORS跨域请求
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func UserLoginHandler(resp http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	uid := query.Get("uid")
	uidInt64, _ := strconv.ParseInt(uid, 10, 64)
}
