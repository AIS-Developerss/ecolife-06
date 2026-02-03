package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// Level представляет уровень логирования
type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelWarn  Level = "WARN"
	LevelError Level = "ERROR"
)

// Logger представляет структурированный логгер
type Logger struct {
	level Level
}

// NewLogger создает новый логгер
func NewLogger(level Level) *Logger {
	return &Logger{level: level}
}

// LogEntry представляет запись лога
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
}

// log записывает лог
func (l *Logger) log(level Level, message string, fields map[string]interface{}) {
	// Проверяем уровень логирования
	if !l.shouldLog(level) {
		return
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     string(level),
		Message:   message,
		Fields:    fields,
	}

	// Форматируем как JSON
	jsonData, err := json.Marshal(entry)
	if err != nil {
		// Fallback на обычный лог
		log.Printf("[%s] %s", level, message)
		return
	}

	// Выводим в stdout (можно перенаправить в файл)
	fmt.Fprintln(os.Stdout, string(jsonData))
}

// shouldLog проверяет, нужно ли логировать на данном уровне
func (l *Logger) shouldLog(level Level) bool {
	levels := map[Level]int{
		LevelDebug: 0,
		LevelInfo:  1,
		LevelWarn:  2,
		LevelError: 3,
	}

	return levels[level] >= levels[l.level]
}

// Debug логирует сообщение уровня DEBUG
func (l *Logger) Debug(message string, fields map[string]interface{}) {
	l.log(LevelDebug, message, fields)
}

// Info логирует сообщение уровня INFO
func (l *Logger) Info(message string, fields map[string]interface{}) {
	l.log(LevelInfo, message, fields)
}

// Warn логирует сообщение уровня WARN
func (l *Logger) Warn(message string, fields map[string]interface{}) {
	l.log(LevelWarn, message, fields)
}

// Error логирует сообщение уровня ERROR
func (l *Logger) Error(message string, fields map[string]interface{}) {
	l.log(LevelError, message, fields)
}

// WithFields создает новый логгер с дополнительными полями
func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	// Для простоты возвращаем тот же логгер
	// В более продвинутых реализациях можно создать обертку
	return l
}

