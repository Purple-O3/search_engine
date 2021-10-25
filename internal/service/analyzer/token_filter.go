package analyzer

import (
	"bufio"
	"io"
	"os"
	"search_engine/internal/util/tools"
)

type tokenFilter struct {
	stopWordDict map[string]bool
}

func newTokenFilter(stopWordPath string) *tokenFilter {
	tf := new(tokenFilter)
	tf.loadStopWordDict(stopWordPath)
	return tf
}

func (tf *tokenFilter) loadStopWordDict(stopWordPath string) {
	tf.stopWordDict = make(map[string]bool)
	fi, err := os.Open(stopWordPath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		term := tools.Bytes2Str(a)
		tf.stopWordDict[term] = true
	}
}

func (tf *tokenFilter) filter(tokens []interface{}) []string {
	terms := make([]string, 0, len(tokens))
	for _, tokenInterface := range tokens {
		token, _ := tokenInterface.(string)
		if _, ok := tf.stopWordDict[token]; !ok {
			terms = append(terms, token)
		}
	}
	return terms
}
