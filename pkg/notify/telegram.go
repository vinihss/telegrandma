package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendTelegram(message string) {

	token := "7247685011:AAFuiFFEVJ--RyY52umhhRNYZ5-0lZZmBEs" // Defina como variável de ambiente
	chatID := "1393047307"                                    // Defina como variável de ambiente

	if token == "" || chatID == "" {
		fmt.Println("Erro: Defina TELEGRAM_BOT_TOKEN e TELEGRAM_CHAT_ID")
		return
	}
	/*cmd := exec.Command("/usr/bin/tgpt", "--provider", "phind", fmt.Sprintf("Criar uma explicação para o log:%s", message))
	if cmd != nil {
		output, err := cmd.Output()
		if err != nil {
			log.Printf("Erro ao enviar alerta: %v", err)
		} else {
			log.Printf("Saída do comando: %s", output)
			message = string(output)
		}

	}*/
	err := sendMessage(token, chatID, message)
	if err != nil {
		fmt.Println("Erro ao enviar mensagem:", err)
	} else {
		fmt.Println("Mensagem enviada com sucesso!")
	}
}

const TELEGRAM_API = "https://api.telegram.org/bot"

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(token, chatID, message string) error {
	url := fmt.Sprintf("%s%s/sendMessage", TELEGRAM_API, token)

	data := TelegramMessage{
		ChatID: chatID,
		Text:   message,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("falha ao enviar mensagem: %s", resp.Status)
	}

	return nil
}
