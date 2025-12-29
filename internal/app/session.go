package app

import (
	"fmt"
	"net"
	"sync"
)

type SessionManager struct {
	sessions sync.Map
}

func NewSessionManager() *SessionManager {
	return &SessionManager{}
}

func (sm *SessionManager) Add(conn net.Conn) {
	addr := conn.RemoteAddr().String()
	sm.sessions.Store(addr, conn)
}

func (sm *SessionManager) Remove(addr string) {
	sm.sessions.Delete(addr)
}

func (sm *SessionManager) Count() int {
	count := 0
	sm.sessions.Range(func(_, _ interface{}) bool {
		count++
		return true
	})
	return count
}

// Gửi tin nhắn tới tất cả client
func (sm *SessionManager) Broadcast(message string) {
	sm.sessions.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			conn.Write([]byte(fmt.Sprintf("[BROADCAST]: %s\n", message)))
		}
		return true
	})
}

// Đóng toàn bộ kết nối
func (sm *SessionManager) CloseAll() {
	sm.sessions.Range(func(key, value interface{}) bool {
		if conn, ok := value.(net.Conn); ok {
			conn.Close()
		}
		return true
	})
}
