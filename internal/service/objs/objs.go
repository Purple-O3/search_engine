package objs

type Posting struct {
	FieldTerm string
	Docid     uint64
	TermFreq  int
	Offset    []int
}

type Postings []Posting    //not same term's posting
type PostingList []Posting //same term's posting

type Data struct {
	Modified string `json:"modified"`
	Saled    string `json:"saled"`
}

type Doc struct {
	Ident string `json:"identification"`
	Data
}

type RecallPosting struct {
	Posting
	Doc
}

type RetreiveTerm struct {
	Field     string `json:"field"`
	FieldData string `json:"fieldData"`
	Operator  string `json:"operator"`
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
		return h[i].Term < h[j].Term
	} else {
		return h[i].Docid < h[j].Docid
	}
}

const (
	Union = "union"
	Inter = "inter"
)
