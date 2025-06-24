package v1

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSHub struct {
	clients map[string]map[*websocket.Conn]bool
	mu      sync.RWMutex
}

func NewWSHub() *WSHub {
	return &WSHub{
		clients: make(map[string]map[*websocket.Conn]bool),
	}
}

func (h *WSHub) Join(boardUUID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.clients[boardUUID] == nil {
		h.clients[boardUUID] = make(map[*websocket.Conn]bool)
	}
	h.clients[boardUUID][conn] = true
}

func (h *WSHub) Leave(boardUUID string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	if h.clients[boardUUID] == nil {
		return
	}
	delete(h.clients[boardUUID], conn)
	if len(h.clients[boardUUID]) == 0 {
		delete(h.clients, boardUUID)
	}
}

func (h *WSHub) Broadcast(boardUUID string, msg any) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for conn := range h.clients[boardUUID] {
		conn.WriteJSON(msg) // FIXME: write message properly
	}
}
