package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocketsからの返却用データの構造体
type WsJsonResponse struct {
	Data string `json:"data"`
}

// WebSocketsコネクション情報を格納
type WebScoketConnection struct {
	*websocket.Conn
}

var (
	// ペイロードチャネルを作成
	wsChan = make(chan WsPayload)

	// コネクションマップを作成
	clients = make(map[WebScoketConnection]string)
)

// WebSockets送信データを格納
type WsPayload struct {
	Data  string              `json:"data"`
	Conn     WebScoketConnection `json:"-"`
}

// wsコネクションの基本設定
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ListenForWs(conn *WebScoketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err == nil {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

// WebSocketsのエンドポイント
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("OK Client Connecting")

	conn := WebScoketConnection{Conn: ws}
	clients[conn] = "user1"

	go ListenForWs(&conn)
}

func broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(&response.Data)
		if err != nil {
			_ = client.Close()
			delete(clients, client)
		}
	}
}

func ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-wsChan
		response.Data = e.Data

		broadcastToAll(response)
	}
}