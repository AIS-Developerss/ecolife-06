package application

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

var (
	// maxStringLength максимальная длина строковых полей
	maxStringLength  = 500
	maxNameLength    = 100
	maxPhoneLength   = 20
	maxAddressLength = 300
)

// ValidatePhone проверяет формат телефона
func ValidatePhone(phone string) error {
	if phone == "" {
		return errors.New("телефон обязателен для заполнения")
	}

	// Удаляем все пробелы, дефисы и скобки для проверки
	cleaned := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(phone, " ", ""), "-", ""), "(", "")
	cleaned = strings.ReplaceAll(strings.ReplaceAll(cleaned, ")", ""), "+", "")

	// Проверяем, что остались только цифры и их количество
	if len(cleaned) < 10 || len(cleaned) > 11 {
		return errors.New("телефон должен содержать 10-11 цифр")
	}

	// Проверяем, что все символы - цифры
	for _, r := range cleaned {
		if r < '0' || r > '9' {
			return errors.New("телефон может содержать только цифры, пробелы, дефисы, скобки и знак +")
		}
	}

	return nil
}

// ValidateName проверяет имя
func ValidateName(name string) error {
	if name == "" {
		return errors.New("имя обязательно для заполнения")
	}

	// Проверяем длину
	if utf8.RuneCountInString(name) > maxNameLength {
		return errors.New("имя слишком длинное (максимум 100 символов)")
	}

	// Проверяем на наличие только пробелов
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return errors.New("имя не может состоять только из пробелов")
	}

	return nil
}

// ValidateStringLength проверяет длину строки
func ValidateStringLength(s string, maxLen int, fieldName string) error {
	if utf8.RuneCountInString(s) > maxLen {
		return errors.New(fieldName + " слишком длинное (максимум " + strconv.Itoa(maxLen) + " символов)")
	}
	return nil
}

// SanitizeString удаляет опасные символы и обрезает пробелы
func SanitizeString(s string) string {
	// Удаляем начальные и конечные пробелы
	s = strings.TrimSpace(s)

	// Удаляем нулевые байты и управляющие символы (кроме пробелов, табуляции, переноса строки)
	var result strings.Builder
	for _, r := range s {
		if r >= 32 || r == '\t' || r == '\n' || r == '\r' {
			result.WriteRune(r)
		}
	}

	return result.String()
}

// ValidateLimitOffset проверяет параметры пагинации
func ValidateLimitOffset(limit, offset int) error {
	if limit < 1 {
		return errors.New("лимит должен быть не менее 1")
	}
	if limit > 100 {
		return errors.New("лимит не может превышать 100")
	}
	if offset < 0 {
		return errors.New("смещение не может быть отрицательным")
	}
	if offset > 10000 {
		return errors.New("смещение не может превышать 10000")
	}
	return nil
}

// ValidateUUID проверяет формат UUID
func ValidateUUID(id string) error {
	uuidRegex := regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
	if !uuidRegex.MatchString(strings.ToLower(id)) {
		return errors.New("неверный формат идентификатора")
	}
	return nil
}
