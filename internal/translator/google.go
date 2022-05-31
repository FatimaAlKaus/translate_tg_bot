package translator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const url = "https://google-translate1.p.rapidapi.com/language/translate/v2"

type GoogleTranslator struct {
	key string
}
type Response struct {
	Data struct {
		Translations []struct {
			TranslatedText string `json:"translatedText"`
		} `json:"translations"`
	} `json:"data"`
}

func (t GoogleTranslator) Translate(source string, target string, q string) (*string, error) {
	s := fmt.Sprintf("source=%s&target=%s&q=%s", source, target, q)
	payload := strings.NewReader(s)
	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", t.key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	text := result.Data.Translations[0].TranslatedText
	return &text, nil
}
func NewGoogleTranslator(key string) *GoogleTranslator {
	return &GoogleTranslator{key: key}
}
