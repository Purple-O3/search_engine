package analyzer

import (
	"search_engine/internal/service/objs"
	"search_engine/internal/util/hashset"
	"search_engine/internal/util/tools"
	"strings"
)

type Analyzer interface {
	Analysis(docid uint64, doc interface{}) objs.Postings
}

func AnalyzerFactory(stopWordPath string) Analyzer {
	return newCustomAnalyzer(stopWordPath)
}

type customAnalyzer struct {
	ti tokenizer
	tf *tokenFilter
}

func newCustomAnalyzer(stopWordPath string) *customAnalyzer {
	ca := new(customAnalyzer)
	ca.ti = tokenizerFactory()
	ca.tf = newTokenFilter(stopWordPath)
	return ca
}

func (ca *customAnalyzer) Analysis(docid uint64, doc interface{}) objs.Postings {
	ps := make(objs.Postings, 0)
	objMap := make(map[string]string, 0)
	objMap, _ = tools.ConvStruct2Map(doc)
	set := hashset.NewSet()

	for fieldName, fieldValue := range objMap {
		set.Clear()
		tokens := ca.ti.getTokens(fieldValue)
		for _, token := range tokens {
			token = strings.TrimSpace(token)
			set.Add(token)
		}
		uniqTokens := set.GetAll()
		terms := ca.tf.filter(uniqTokens)

		for _, term := range terms {
			var posting objs.Posting
			posting.FieldTerm = fieldName + "_" + term
			posting.Docid = docid
			/*posting.TermFreq = strings.Count(doc, term)
			posting.Offset = make([]int, 0)
			offset := 0
			docStart := 0
			//TODO:需优化
			for i := 0; i < posting.TermFreq; i++ {
				offset = docStart + strings.Index(doc[docStart:], term)
				posting.Offset = append(posting.Offset, offset)
				docStart = offset + len(term)
			}*/
			ps = append(ps, posting)
		}
	}
	return ps
}
