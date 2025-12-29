package app

import (
	"context"
	"fmt"
	"net"
	"sync"
)

type App struct {
	Listener net.Listener
	Sessions *SessionManager
	wg       sync.WaitGroup
}

func NewApp(l net.Listener, sm *SessionManager) *App {
	return &App{
		Listener: l,
		Sessions: sm,
	}
}

func (a *App) Start(ctx context.Context) {
	for {
		conn, err := a.Listener.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Printf("Accept error: %v\n", err)
				continue
			}
		}

		a.wg.Add(1)
		go a.handleConnection(ctx, conn)
	}
}

func (a *App) handleConnection(ctx context.Context, conn net.Conn) {
	defer a.wg.Done()
	addr := conn.RemoteAddr().String()

	a.Sessions.Add(conn)
	fmt.Printf("[+] Session opened: %s\n", addr)

	defer func() {
		a.Sessions.Remove(addr)
		conn.Close()
		fmt.Printf("[-] Session closed: %s\n", addr)
	}()

	// Tạo buffer để đọc dữ liệu
	buf := make([]byte, 4096)

	for {
		select {
		case <-ctx.Done():
			conn.Write([]byte("Server shutting down...\n"))
			return
		default:
			n, err := conn.Read(buf)
			if err != nil {
				return
			}

			// Logic processing
			message := string(buf[:n])
			fmt.Printf("[%s]: %s", addr, message)

			// Send response but do not close
			conn.Write([]byte("ACK: Received your message\n"))
		}
	}
}

func (a *App) Stop() {
	fmt.Println("Shutting down...")
	a.Listener.Close()
	a.Sessions.CloseAll()
	a.wg.Wait()
	fmt.Println("Server stopped cleanly.")
}
