/*
Defines the models used in this web application as structs.
Hassaan Ali Wattoo <hawattoo@umich.edu>
*/

package model

import (
	"time"
)

type Word struct {
	DatePublished    time.Time `json:"date_published"`
	WordRomanUrdu    string    `json:"word_roman_urdu"`
	WordNastaliqUrdu string    `json:"word_nastaliq_urdu"`
	WordMeaning      string    `json:"word_meaning"`

	// Each Word is associated with a verse from a Ghazal for further meaning
	// Here we store the metadata for that Ghazal
	FirstMisra             string `json:"first_misra"`
	SecondMisra            string `json:"second_misra"`
	FirstMisraTranslation  string `json:"first_misra_translation"`
	SecondMisraTranslation string `json:"second_misra_translation"`
	GhazalName             string `json:"ghazal_name"`
	GhazalNameLink         string `json:"ghazal_name_link"`
	GhazalAuthor           string `json:"ghazal_author"`
	GhazalAuthorLink       string `json:"ghazal_author_link"`
}
