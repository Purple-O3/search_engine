package objs

type AddDocResp struct {
	Docid uint64 `json:"docid"`
}

type RetreiveDocResp struct {
	Count  int               `json:"count"`
	Result RecallPostingList `json:"result""`
}
