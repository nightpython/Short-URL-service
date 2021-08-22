package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "e0dba740-fc4b-4977-872c-d360239e6b1a"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://yandex.ru/"
	shortLink_1 := UrlGenerator(initialLink_1, UserId)

	initialLink_2 := "https://en.wikipedia.org/wiki/Main_Page"
	shortLink_2 := UrlGenerator(initialLink_2, UserId)

	initialLink_3 := "https://github.com/nightpython"
	shortLink_3 := UrlGenerator(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "4J67wvBh")
	assert.Equal(t, shortLink_2, "wiZFZhsu")
	assert.Equal(t, shortLink_3, "i07uKW0e")
}
