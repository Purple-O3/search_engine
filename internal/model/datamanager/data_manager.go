package datamanager

import (
	"encoding/json"
	"search_engine/internal/model/index"
	"search_engine/internal/model/store"
	"search_engine/internal/objs"
	"search_engine/internal/util/log"
	"search_engine/internal/util/tools"
	"strconv"
	"time"
)

type Manager struct {
	invertedBuffer *index.InvertedIndex
	positiveBuffer *index.PositiveIndex
	db             store.Store
}

func NewManager(config objs.DBConfig) *Manager {
	mg := new(Manager)
	var err error
	mg.db, err = store.StoreFactory(config)
	if err != nil {
		panic(err)
	}
	mg.invertedBuffer = index.NewInvertedIndex(mg.db)
	mg.positiveBuffer = index.NewPositiveIndex(mg.db)
	return mg
}

func (mg *Manager) Close() {
	mg.invertedBuffer.FlushAll()
	mg.positiveBuffer.FlushAll()
	mg.db.Close()
}

func (mg *Manager) getDbPl(k string) (objs.PostingList, error) {
	key := tools.Str2Bytes(k)
	value, err := mg.db.Get(key)
	if err != nil {
		return nil, err
	}
	pl := make(objs.PostingList, 0)
	if err := json.Unmarshal(value, &pl); err != nil {
		return nil, err
	}
	return pl, nil
}

func (mg *Manager) getDbStr(k string) (string, error) {
	key := tools.Str2Bytes(k)
	value, err := mg.db.Get(key)
	if err != nil {
		return "", err
	}
	v := tools.Bytes2Str(value)
	return v, nil
}

func (mg *Manager) AddDoc(doc objs.Doc, docid uint64, ps objs.Postings) {
	for _, posting := range ps {
		mg.invertedBuffer.Set(posting.FieldName+"_"+posting.Term, posting)
	}
	docidString := strconv.FormatUint(docid, 10)
	docKey := "doc" + docidString
	docByte, _ := json.Marshal(doc)
	mg.positiveBuffer.Set(docKey, tools.Bytes2Str(docByte))
}

func (mg *Manager) Retrieve(fieldName string, term string, trackid uint64) (objs.RecallPostingList, error) {
	defer func(cost func() time.Duration) {
		log.Warnf("trackid:%v, cost: %.3f ms", trackid, float64(cost().Microseconds())/1000.0)
	}(tools.TimeCost())

	pl := make(objs.PostingList, 0)
	fieldKey := fieldName + "_" + term
	if plBuffer, ok := mg.invertedBuffer.Get(fieldKey); ok {
		pl = append(pl, plBuffer...)
	}
	if plDb, err := mg.getDbPl(fieldKey); err == nil {
		pl = append(pl, plDb...)
	}

	recallPl := make(objs.RecallPostingList, 0, len(pl))
	for _, posting := range pl {
		var docString string
		docidString := strconv.FormatUint(posting.Docid, 10)
		docKey := "doc" + docidString
		if valueBuffer, ok := mg.positiveBuffer.Get(docKey); ok {
			docString = valueBuffer
		} else {
			if valueDb, err := mg.getDbStr(docKey); err == nil {
				docString = valueDb
			}
		}

		postingRec := objs.RecallPosting{}
		postingRec.Posting = posting
		if err := json.Unmarshal(tools.Str2Bytes(docString), &postingRec.Doc); err != nil {
			return nil, err
		}
		recallPl = append(recallPl, postingRec)
	}

	log.Debugf("trackid:%v, repl:%v", trackid, recallPl)
	return recallPl, nil
}
