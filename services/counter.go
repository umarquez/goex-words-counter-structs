package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type Counter struct {
	fullText      string
	words         []string
	wordsCounters map[string]int
}

type TxtGenerator func() (string, error)

func NewCounter(generator TxtGenerator) (Counter, error) {
	text, err := generator()
	if err != nil {
		err = fmt.Errorf(
			"text generator fails: %v", err,
		)
		return Counter{}, err
	}

	counter := Counter{
		fullText: text,
	}
	counter.init()

	return counter, nil
}

func (c *Counter) init() error {
	err := c.extractWords()
	if err != nil {
		err = fmt.Errorf(
			"text generator fails: %v", err,
		)
		return err
	}

	c.countWordsOccurrences()

	return nil
}

func (c *Counter) extractWords() error {
	c.words = regexp.MustCompile(`[a-zA-Z/-]+`).FindAllString(c.fullText, -1)
	return nil
}

func (c *Counter) countWordsOccurrences() {
	c.wordsCounters = make(map[string]int)
	for _, word := range c.words {
		c.wordsCounters[strings.ToLower(word)]++
	}

	return
}

func (c Counter) GetWordsList() []string {
	var words []string
	for key, _ := range c.wordsCounters {
		words = append(words, key)
	}

	return words
}

func (c Counter) GetWordTotal(word string) (int, error) {
	total, ok := c.wordsCounters[word]
	if !ok {
		return 0, errors.New("this word was not found")
	}
	return total, nil
}

func (c Counter) String() string {
	buffJson := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(buffJson)
	jsonEncoder.SetIndent("", "  ")
	err := jsonEncoder.Encode(c.wordsCounters)
	if err != nil {
		log.Fatal(fmt.Errorf("impossible to encode counters: %v", err))
	}
	return fmt.Sprintf("Original text:%v\n\n%v", c.fullText, buffJson.String())
}
