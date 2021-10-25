package rank

import (
	"search_engine/internal/service/objs"
	"testing"
)

func TestAll(t *testing.T) {
	ranker := RankerFactory()
	replUniqInter := make(objs.RecallPostingList, 0)
	reposting := objs.RecallPosting{objs.Posting{Term: "银河", Docid: 1, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五二班", Price: 5.200000}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 2, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{objs.Posting{Term: "银河", Docid: 4, TermFreq: 1, Offset: []int{0}}, objs.Doc{Body: "", Title: "五三班", Price: 5.300000}}
	replUniqInter = append(replUniqInter, reposting)
	out := ranker.Rank(replUniqInter)
	t.Log(out)
}
