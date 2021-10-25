package analyzer

import (
	"github.com/yanyiwu/gojieba"
)

type tokenizer interface {
	getTokens(doc string) []string
}

func tokenizerFactory() tokenizer {
	return newJieBa()
}

type jieBa struct {
	jb *gojieba.Jieba
}

func newJieBa() *jieBa {
	j := new(jieBa)
	j.jb = gojieba.NewJieba()
	return j
}

func (j *jieBa) getTokens(doc string) []string {
	useHmm := true
	tokens := j.jb.CutForSearch(doc, useHmm)
	return tokens
}
