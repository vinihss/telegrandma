package notify

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

// Envia um alerta quando um recurso está crítico
func SendAlert(message string) {
	log.Printf("[ALERTA] %s", message)

	switch runtime.GOOS {
	case "linux":
		exec.Command("notify-send", "Alerta do Monitor", message).Run()
	case "darwin":
		exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "Alerta do Monitor"`, message)).Run()
	case "windows":
		exec.Command("powershell", "-Command", `New-BurntToastNotification -Text "Alerta do Monitor", "`+message+`"`).Run()
	}

}
