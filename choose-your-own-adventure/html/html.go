package html

import (
	"encoding/json"
	"log"
)

type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func Unmarshaller(text []byte, ch map[string]Chapter) {

	err := json.Unmarshal(text, &ch)
	if err != nil {
		log.Fatal(err)
	}
}
