package test

import (
	"github.com/hassaanaliw/seekhlai-api/scraper"
	"testing"
	"time"
)

func TestScrape(t *testing.T) {
	scraper.ScrapeTodayWord(time.Now())
}
