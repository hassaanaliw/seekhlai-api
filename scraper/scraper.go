package scraper

import (
	"fmt"
	"time"
	"github.com/anaskhan96/soup"
	"os"
	"github.com/hassaanaliw/seekhlai-api/model"
)

// Function that loads the word of the day page from rekhta.org for the
// current day and scrapes it for the appropriate information
// Rekhta URL in the format of https://rekhta.org/archives/YEAR-MONTH-DAY/wordofthedays/
// Returns a Word Object
func ScrapeTodayWord(date time.Time) model.Word {

	// Format todays date in the format YEAR-MONTH-DATE
	currentTime := date
	currentDate := fmt.Sprintf("%d-%d-%d", currentTime.Year(), currentTime.Month(), currentTime.Day())

	word := model.Word{}
	word.DatePublished = currentTime

	// request and parse the the word of the day page for today
	baseURL := "https://rekhta.org/archives/%s/wordofthedays/"
	resp, err := soup.Get(fmt.Sprintf(baseURL, currentDate))

	if err != nil {
		fmt.Println(err)
        os.Exit(1)
	}

	doc := soup.HTMLParse(resp)

	word.WordRomanUrdu = doc.Find("div", "class", "wordContainer").Find("h1").Text()
	word.WordNastaliqUrdu = doc.Find("li", "class", "urMeaning").Text()
	word.WordMeaning = doc.Find("div", "class", "engMeaning").Find("h3").Text()
	word.GhazalAuthor = doc.Find("div", "class", "sherDetail").Find("a").Text()
	word.GhazalAuthorLink = "https://rekhta.org" +
		doc.Find("div", "class", "sherDetail").Find("a").Attrs()["href"]

	verseOne := ExtractVerseAndTranslation(doc, 1)
	word.FirstMisra = verseOne[0]
	word.FirstMisraTranslation = verseOne[1]

	verseTwo := ExtractVerseAndTranslation(doc, 2)
	word.SecondMisra = verseTwo[0]
	word.SecondMisraTranslation = verseTwo[1]

	return word

}

// Given a Root document and the number of the verse to extract, returns
// a slice of the verse and it's translation
func ExtractVerseAndTranslation(doc soup.Root, number int) []string {
	parentDiv := doc.Find("div", "class", "c")
	// Select the paragraph based on the number passed to this function
	paragraph := parentDiv.FindAll("p")[number-1]

	// Iterate over all the paragraph spans and join their text together to form the verse
	spans := paragraph.FindAll("span")
	verseText := ""
	for _, span := range spans {
		verseText += span.Text()
		verseText += " "
	}

	translationDiv := doc.Find("div", "class", "t")
	if (translationDiv.Pointer == nil) {
		// Some days, the rekhta page does not include a translation for the verses
        // Return empty string in that case
        return []string{verseText, " "}
	}
	translationParagraph := translationDiv.FindAll("p")[number-1]

	// Iterate over all the paragraph spans and join their text together to form the translations
	spans = translationParagraph.FindAll("span")
	translationText := ""
	for _, span := range spans {
		translationText += span.Text()
		translationText += " "
	}

	return []string{verseText, translationText}
}
