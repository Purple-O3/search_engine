package objs

import "time"

type Posting struct {
	FieldName string `json:"-"`
	Term      string `json:"-"`
	Docid     uint64
	//TermFreq  int
	//Offset    []int
}

type Postings []Posting    //not same term's posting
type PostingList []Posting //same term's posting

type Data struct {
	Modified  string
	Saled     string
	Num       int
	CreatedAt time.Time `search_type:"keyword"`
}

type Doc struct {
	Ident string `search_type:"keyword"`
	Data
}

type RecallPosting struct {
	Posting
	Doc
}

type FieldInfo struct {
	Type  string
	Value string
}

const NotSplit = "keyword"

const (
	Union  = "should"
	Inter  = "must"
	Filter = "filter"
)

const (
	Eq  = 0x001
	Gt  = 0x010
	Gte = 0x011
	Lt  = 0x100
	Lte = 0x101
)

type RetreiveTerm struct {
	FieldName       string
	Term            interface{}
	TermCompareType int
	Operator        string
}

type RetreiveReq struct {
	RetreiveTerms []RetreiveTerm
	Offset        int
	Limit         int
}

type RecallPostingList []RecallPosting

//实现排序
func (h RecallPostingList) Len() int {
	return len(h)
}

func (h RecallPostingList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h RecallPostingList) Less(i, j int) bool {
	if h[i].Docid == h[j].Docid {
		if h[i].FieldName == h[j].FieldName {
			return h[i].Term < h[j].Term
		} else {
			return h[i].FieldName < h[j].FieldName
		}
	} else {
		return h[i].Docid < h[j].Docid
	}
}
