package translator

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	text        = "{TEXT}"
	from        = "{FROM}"
	to          = "{TO}"
	myMemoryurl = "https://translated-mymemory---translation-memory.p.rapidapi.com/api/get?q={TEXT}&langpair={FROM}%7C{TO}"
)

type MyMemoryTranslator struct {
	key string
}

func (t MyMemoryTranslator) Translate(source string, target string, q string) (*string, error) {
	q = strings.Join(strings.Split(q, " "), "%20")
	url := strings.Replace(strings.Replace(strings.Replace(myMemoryurl, text, q, 1), from, source, 1), to, target, 1)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-RapidAPI-Host", "translated-mymemory---translation-memory.p.rapidapi.com")
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
	var result map[string]interface{} = make(map[string]interface{})
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	text := result["responseData"].(map[string]interface{})["translatedText"].(string)
	return &text, nil
}
func NewMeMemoryTranslator(key string) *MyMemoryTranslator {
	return &MyMemoryTranslator{key: key}
}
