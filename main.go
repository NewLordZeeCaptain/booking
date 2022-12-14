package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Name, Description, URL string
}

type SearchResults struct {
	ready  bool
	Query  string
	Result []Result
}

func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Result = append(sr.Result, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

// This func will get and return result from Wikipedia
func wikipediaAPI(request string) (answer []string) {
	// Slice of 3 elements
	s := make([]string, 3)

	//Send request
	if response, err := http.Get(request); err != nil {
		s[0] = "Wikipedia is not responds"
	} else {
		defer response.Body.Close()

		//Read answer
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		sr := &SearchResults{}
		if err = json.Unmarshal([]byte(contents), sr); err != nil {
			s[0] = "Something going wrong, try to change your question"
		}

		if !sr.ready {
			s[0] = "Something going wrong, try to change your question"
		}

		for i := range sr.Result {
			s[i] = sr.Result[i].URL
		}
	}
	return s

}

func main() {

}
