package utils

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"github.com/micmonay/keybd_event"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	// Inicializando PortAudio
	err := portaudio.Initialize()
	if err != nil {
		fmt.Println("Erro ao inicializar PortAudio:", err)
		return
	}
	defer portaudio.Terminate()

	// Configurando parâmetros de gravação
	buffer := make([]int16, 44100) // 1 segundo de áudio com 44100 Hz
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(buffer), buffer)
	if err != nil {
		fmt.Println("Erro ao abrir o stream:", err)
		return
	}
	defer stream.Close()

	// Configurando o socket
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println("Erro ao conectar ao socket:", err)
		return
	}
	defer conn.Close()

	// Configurando o capturador de teclas
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("Erro ao inicializar o capturador de teclas:", err)
		return
	}
	kb.SetKeys(keybd_event.VK_R)
	kb.HasCTRL(true)
	kb.HasALT(true)

	// Canal para capturar sinais de interrupção
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Daemon iniciado. Pressione Ctrl+Alt+R para começar a gravar.")

	// Loop principal
	for {
		select {
		case <-sigChan:
			fmt.Println("Encerrando daemon...")
			return
		default:
			err := kb.Launching()
			if err == nil {
				fmt.Println("Iniciando gravação...")
				err = stream.Start()
				if err != nil {
					fmt.Println("Erro ao iniciar o stream:", err)
					return
				}

				// Grava e envia áudio enquanto a tecla estiver pressionada
				err := kb.Launching()
				for err == nil {
					err = stream.Read()
					if err != nil {
						fmt.Println("Erro durante a leitura do áudio:", err)
						return
					}
					_, err = conn.Write(toByteSlice(buffer))
					if err != nil {
						fmt.Println("Erro ao enviar áudio para o socket:", err)
						return
					}
				}

				err = stream.Stop()
				if err != nil {
					fmt.Println("Erro ao parar o stream:", err)
					return
				}
				fmt.Println("Gravação finalizada.")
			}
		}
	}
}

// Função auxiliar para converter []int16 para []byte
func toByteSlice(data []int16) []byte {
	byteData := make([]byte, len(data)*2)
	for i, v := range data {
		byteData[i*2] = byte(v)
		byteData[i*2+1] = byte(v >> 8)
	}
	return byteData
}
