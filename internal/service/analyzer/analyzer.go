package analyzer

import (
	"search_engine/internal/service/objs"
	"search_engine/internal/util/hashset"
	"search_engine/internal/util/tools"
	"strings"
)

const NotSplit = "keyword"

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
	//TODO:挪位置
	fieldMap, _ := tools.ConvStruct2Map(doc)
	set := hashset.NewSet()

	for fieldName, fieldInfo := range fieldMap {
		fieldValue := fieldInfo.Value
		fieldType := fieldInfo.Type
		var terms []string
		if fieldType == NotSplit {
			terms = []string{fieldValue}
		} else {
			set.Clear()
			tokens := ca.ti.getTokens(fieldValue)
			for _, token := range tokens {
				token = strings.TrimSpace(token)
				set.Add(token)
			}
			uniqTokens := set.GetAll()
			terms = ca.tf.filter(uniqTokens)
		}

		for _, term := range terms {
			var posting objs.Posting
			posting.FieldName = fieldName
			posting.Term = term
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
