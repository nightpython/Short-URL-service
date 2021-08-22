package shortener

import (
	"crypto/sha256"
	"encoding/base64"
)

//функция для хэширования и кодировки исходных данных, складываем userID и input для уникальности ссылки
func UrlGenerator(input string, userId string) string {
	algorithm := sha256.New()
	algorithm.Write([]byte(input + userId))
	encoded := base64.StdEncoding.EncodeToString(algorithm.Sum(nil))
	return encoded[:8]
}
