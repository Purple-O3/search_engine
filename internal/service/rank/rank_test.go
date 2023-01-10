package rank

import (
	"search_engine/internal/objs"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	ranker := RankerFactory()
	replUniqInter := make(objs.RecallPostingList, 0)
	reposting := objs.RecallPosting{Posting: objs.Posting{FieldName: "Modified", Term: "银河", Docid: 1}, Doc: objs.Doc{Ident: "88.199.1/aaa.def", Data: objs.Data{Modified: "北京市石景山区", Saled: "乌鲁木齐", CreatedAt: time.Now().Add(time.Hour * 24), Num: 15}}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{Posting: objs.Posting{FieldName: "Modified", Term: "银河", Docid: 2}, Doc: objs.Doc{Ident: "88.199.1/bbb.def", Data: objs.Data{Modified: "北京市丰台区", Saled: "辽宁", CreatedAt: time.Now().Add(time.Second * 1), Num: 13}}}
	replUniqInter = append(replUniqInter, reposting)
	reposting = objs.RecallPosting{Posting: objs.Posting{FieldName: "Modified", Term: "银河", Docid: 4}, Doc: objs.Doc{Ident: "88.199.1/ccc.def", Data: objs.Data{Modified: "北京市宣武区", Saled: "大连", CreatedAt: time.Now().Add(time.Hour * 6), Num: 10}}}
	replUniqInter = append(replUniqInter, reposting)
	out := ranker.Rank(replUniqInter)
	t.Log(out)
}
