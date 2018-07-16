package scraper

import (
	"fmt"
	"os"
	"time"

	"github.com/anaskhan96/soup"
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

	ExtractGhazalAuthorAndName(doc, &word)

	verseOne := ExtractVerseAndTranslation(doc, 1)
	word.FirstMisra = verseOne[0]
	word.FirstMisraTranslation = verseOne[1]

	verseTwo := ExtractVerseAndTranslation(doc, 2)
	word.SecondMisra = verseTwo[0]
	word.SecondMisraTranslation = verseTwo[1]

	return word

}

// ExtractVerseAndTranslation , given a Root document and the number of the verse to extract, returns
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
	if translationDiv.Pointer == nil {
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

// ExtractGhazalAuthorAndName given a doc and a word, extracts either both the author and ghazal
// details or just the author info based on what the Rekhta page contains and adds these to the
// word object. It returns a reference to the word.
func ExtractGhazalAuthorAndName(doc soup.Root, word *model.Word) *model.Word {
	mainDivLinks := doc.Find("div", "class", "sherDetail").FindAll("a")
	if len(mainDivLinks) > 1 {
        // If only one link is returned, this means that Rekhta only provided the Author 
        // information for that ghazal, not a link to the ghazal itself. 
        word.GhazalName = mainDivLinks[0].Text()
		word.GhazalNameLink = "https://rekhta.org" + mainDivLinks[0].Attrs()["href"]
		word.GhazalAuthor = mainDivLinks[1].Text()
		word.GhazalAuthorLink = "https://rekhta.org" + mainDivLinks[1].Attrs()["href"]
	} else {
		word.GhazalAuthor = mainDivLinks[0].Text()
		word.GhazalAuthorLink = "https://rekhta.org" + mainDivLinks[0].Attrs()["href"]
		word.GhazalName = " "
		word.GhazalNameLink = " "
	}

	return word

}
