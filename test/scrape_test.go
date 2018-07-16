package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hassaanaliw/seekhlai-api/scraper"
	"github.com/stretchr/testify/assert"
)

func TestScrape(t *testing.T) {
	testDate, error := time.Parse("2006-01-02", "2018-07-13")
	if error != nil {
		panic(error)
	}
	// Fetch the JSON for a date for which we know what the JSON response should be
	wordJSON := scraper.ScrapeTodayWord(testDate)

	assert.Equal(t, wordJSON.WordRomanUrdu, "ziyaa.n")
	assert.Equal(t, wordJSON.WordNastaliqUrdu, "زیاں")
	assert.Equal(t, wordJSON.WordMeaning, "loss")
	assert.Equal(t, wordJSON.FirstMisra, "ai  dil  tamām  naf.a  hai  saudā-e-ishq  meñ  ")
	assert.Equal(t, wordJSON.SecondMisra, "ik  jaan  kā  ziyāñ  hai  so  aisā  ziyāñ  nahīñ  ")
	assert.Equal(t, wordJSON.FirstMisraTranslation, " ")
	assert.Equal(t, wordJSON.SecondMisraTranslation, " ")
	assert.Equal(t, wordJSON.GhazalName, " ")
	assert.Equal(t, wordJSON.GhazalNameLink, " ")
	assert.Equal(t, wordJSON.GhazalAuthorLink, "https://rekhta.org/poets/mufti-sadruddin-aazurda")
	assert.Equal(t, wordJSON.GhazalAuthor, "Mufti Sadruddin Aazurda")

}

func TestScrapeToday(t *testing.T) {
	testDate := time.Now()
	// Fetch the JSON for a date for which we know what the JSON response should be
	word := scraper.ScrapeTodayWord(testDate)
	fmt.Println(word)

}
