package logger

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

// Consulta os últimos X logs no banco de dados
func GetRecentLogs(limit int) ([]Log, error) {
	var logs []Log
	result := db.Order("created_at desc").Limit(limit).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

// Insere um novo log no banco de dados
func Create(level string, message string) {
	fmt.Println("Create logger method")
	logEntry := Log{
		Message: message,
		Level:   level,
	}
	result := db.Create(&logEntry)
	if result.Error != nil {
		log.Fatal("Erro ao inserir log:", result.Error)
	}
	fmt.Println("Log inserido com sucesso! ID:", logEntry.ID)
}

// Consulta todos os logs no banco de dados
func getLogs(db *gorm.DB) ([]Log, error) {
	var logs []Log
	result := db.Order("created_at desc").Find(&logs) // Ordena por data de criação (do mais recente para o mais antigo)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func init() {
	log.Printf("init logger")

	// Conecta ao banco de dados SQLite
	var err error
	db, err = gorm.Open(sqlite.Open("logs.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}

	// Configura o banco de dados (cria a tabela automaticamente)
	err = db.AutoMigrate(&Log{})
	if err != nil {
		log.Fatal("Erro ao migrar o banco de dados:", err)
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// Estrutura para representar um log
type Log struct {
	gorm.Model        // Adiciona campos padrão como ID, CreatedAt, UpdatedAt, DeletedAt
	Message    string `gorm:"size:255"` // Mensagem do log
	Level      string `gorm:"size:50"`  // Nível de severidade (INFO, ERROR, DEBUG, etc.)
}

func CreateLog(message string, level string) {
	fmt.Println("Create logger method")
	logEntry := Log{
		Message: message,
		Level:   level,
	}

	result := db.Create(&logEntry)
	if result.Error != nil {
		log.Fatal("Erro ao inserir log:", result.Error)
	}
	fmt.Println("Log inserido com sucesso! ID:", logEntry.ID)
}
