package ipc

import (
	"copybara/notifications"
	"copybara/utils"
	"sync"

	ipc "github.com/james-barrow/golang-ipc"
)

const socketName = "copybaraclipboardautomationsocket"

type SafeBool struct {
	mu sync.Mutex
	v  bool
}

var IsCopybaraEnabled = SafeBool{v: true}

func (c *SafeBool) Toggle() {
	c.mu.Lock()
	c.v = !c.v
	c.mu.Unlock()
}

func (c *SafeBool) Value() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

func Init() {
	s, err := ipc.StartServer(socketName, nil)
	utils.CheckError(err)

	for {
		message, _ := s.Read()
		if string(message.Data) == "t" {
			s.Write(1, []byte("t"))
			IsCopybaraEnabled.Toggle()
			if IsCopybaraEnabled.Value() {
				notifications.SendNotification("Toggled on", "list-add")
			} else {
				notifications.SendNotification("Toggled off", "list-remove")
			}
		}
	}

}

func SendToggleCommand() {
	c, err := ipc.StartClient(socketName, nil)
	utils.CheckError(err)
	for {
		message, err := c.Read()
		utils.CheckError(err)
		if message.MsgType == -1 {
			if c.Status() == "Connected" {
				err = c.Write(1, []byte("t"))
				utils.CheckError(err)
			}
		} else {
			if string(message.Data) == "t" {
				c.Close()
				return
			}
		}

	}

}
