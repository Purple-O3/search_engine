package customnet

import "search_engine/internal/service/objs"

type RetreiveReq struct {
	RetreiveTerms []objs.RetreiveTerm
}

type ResultData struct {
	Repl  objs.RecallPostingList `json:"repl"`
	Docid uint64                 `json:"docid"`
}

type RespData struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Result  ResultData `json:"result"`
}
