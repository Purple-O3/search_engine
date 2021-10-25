package tools

import "testing"

func TestAll(t *testing.T) {
	type Data struct {
		Modified string `json:"modified"`
		Saled    string `json:"saled"`
	}
	type Doc struct {
		Ident string `json:"identification"`
		Data
	}

	doc := Doc{Ident: "88.199.1/abc.def", Data: Data{Modified: "北京", Saled: "北京"}}
	objMap, err := ConvStruct2Map(doc)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(objMap)
}
