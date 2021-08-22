package store

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializedStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestPutAndGet(t *testing.T) {
	initialLink := "https://www.yandex.ru"
	shortURL := "Asz4k57oAX"

	SaveUrlMapping(shortURL, initialLink)

	revivedUrl := GetInitialUrl(shortURL)

	fmt.Println(assert.Equal(t, initialLink, revivedUrl))
}
