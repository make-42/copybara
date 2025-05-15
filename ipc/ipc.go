package ipc

import (
	"context"
	"copybara/notifications"
	"copybara/utils"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

const socketName = "/tmp/copybaraclipboardautomationsocket.sock"

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
	router := gin.Default()

	router.GET("/toggle", func(c *gin.Context) {
		IsCopybaraEnabled.Toggle()
		if IsCopybaraEnabled.Value() {
			notifications.SendNotification("Toggled on", "list-add")
		} else {
			notifications.SendNotification("Toggled off", "list-remove")
		}
		c.String(http.StatusOK, fmt.Sprintf("OK"))
	})

	listener, err := net.Listen("unix", socketName)
	utils.CheckError(err)

	http.Serve(listener, router)
}

func SendToggleCommand() {
	conn, err := net.Dial("unix", socketName)
	utils.CheckError(err)

	client := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return conn, nil
			},
		},
	}

	resp, err := client.Get("http://unix/toggle")
	utils.CheckError(err)
	resp.Body.Close()
}
