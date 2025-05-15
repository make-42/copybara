package listener

import (
	"bufio"
	"copybara/config"
	"copybara/ipc"
	"copybara/notifications"
	"copybara/regex"
	"copybara/urlclean"
	"copybara/utils"
	"fmt"
	"os/exec"
	"sync"
	"time"
)

var ListenerInterval = 100 * (time.Millisecond)

var OldText = SafeText{}

type SafeText struct {
	mu   sync.Mutex
	text string
}

func (c *SafeText) Set(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access c.text.
	c.text = key
	c.mu.Unlock()
}

func (c *SafeText) Value() string {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the c.text.
	defer c.mu.Unlock()
	return c.text
}

func ListenerThread() {
	cmd := exec.Command("wl-paste", "-w", "echo", "new")

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if ipc.IsCopybaraEnabled.Value() {
			textBytes, err := exec.Command("wl-paste").Output()
			text := string(textBytes)
			text = text[:len(text)-1] // Strip last character that corresponds to a newline
			utils.CheckError(err)
			if text != OldText.Value() {
				OldText.Set(text)
				newText := text
				urlCleaned := false
				regexReplaced := false
				if config.Config.EnableURLCleaning {
					newText, urlCleaned = urlclean.CleanURLs(newText)
				}
				if config.Config.EnableRegexAutomations {
					newText, regexReplaced = regex.Clean(newText)
				}
				OldText.Set(newText)
				if urlCleaned || regexReplaced {
					err := exec.Command("wl-copy", newText).Run()
					utils.CheckError(err)
					if config.Config.NotificationsOnAppliedAutomations {
						notificationText := fmt.Sprintf("Automations applied to copied text:\n\n[%s]\n->[%s]\n\n", text, newText)
						if urlCleaned {
							notificationText += "[URL]"
						}
						if regexReplaced {
							notificationText += "[REGEX]"
						}
						notifications.SendNotification(notificationText, "edit-find-replace")
					}
				}
			}
		}
	}
}
