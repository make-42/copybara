package notifications

import notify "github.com/TheCreeper/go-notify"

func SendNotification(body string, icon string) {
	ntf := notify.NewNotification("Copybara ₍ᐢ•(ܫ)•ᐢ₎", body)
	ntf.AppIcon = icon
	ntf.Show()
}
