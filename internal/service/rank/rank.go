package rank

import (
	"search_engine/internal/objs"
)

type Ranker interface {
	Rank(in objs.RecallPostingList) objs.RecallPostingList
}

func RankerFactory() Ranker {
	return newVectorSpaceModel()
}

type vectorSpaceModel struct {
}

func newVectorSpaceModel() *vectorSpaceModel {
	vs := new(vectorSpaceModel)
	return vs
}

func (vs *vectorSpaceModel) Rank(in objs.RecallPostingList) objs.RecallPostingList {
	return in
}
