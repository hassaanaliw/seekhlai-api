package test

import (
	"testing"
	"github.com/hassaanaliw/seekhlai-api/scraper"
	"time"
)

func TestScrape(t *testing.T) {
	scraper.ScrapeTodayWord(time.Now())
}
