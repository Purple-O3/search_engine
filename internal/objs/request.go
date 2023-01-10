package objs

type DocidReq struct {
	Docid   uint64 `form:"docid"`
	Trackid uint64 `header:"X-Trackid"`
}

type DocReq struct {
	Doc
	Trackid uint64 `header:"X-Trackid"`
}

type RetreiveReq struct {
	RetreiveTerms []RetreiveTerm
	Offset        int
	Limit         int
	Trackid       uint64 `header:"X-Trackid"`
}
