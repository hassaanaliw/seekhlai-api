# Seekh Lai API

[![Build Status](https://travis-ci.com/hassaanaliw/seekhlai-api.svg?branch=master)](https://travis-ci.com/hassaanaliw/seekhlai-api)

An API for the Rektha.org Word of the Day function in GoLang. 

It is the API behind the Seekh Lai Chrome Extension I'm developing as a personal language learning project.


# Usage

First install the package 

```bash
go get github.com/hassaanaliw/seekhlai-api
```

```go

package main

import (
	"github.com/hassaanaliw/seekhlai-api/scraper"
	"time"
)

func main() {
  # To fetch the word for any other day, pass the time object for that date
	scraper.ScrapeTodayWord(time.Now())
}

```

# Sample API Response

``` Javascript

{
  date_published: "2018-07-02T10:38:55.060461898+05:00",
  word_roman_urdu: "lahad",
  word_nastaliq_urdu: "لحد",
  word_meaning: "burial chamber / grave",
  first_misra: "lahad meñ kyuuñ na jā.ūñ muñh chhupā.e ",
  second_misra: "bharī mahfil se uThvāyā gayā huuñ ",
  first_misra_translation: "why should I not be interred with a covered face ",
  second_misra_translation: "I have been cast from her presence in such disgrace ",
  ghazal_name: "",
  ghazal_name_link: "",
  ghazal_author: "Shad Azimabadi",
  ghazal_author_link: "https://rekhta.org/poets/shad-azimabadi"
}

```

Note: Only some days have the ghazal name and link included in the API response.



