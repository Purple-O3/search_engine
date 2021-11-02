package analyzer

import (
	"search_engine/internal/service/objs"
	"testing"
	"time"
)

/*func TestOffset(t *testing.T) {
	t.Log("test1")
	stopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	analyzer := AnalyzerFactory(stopWordPath)
	var docid uint64 = 0
	doc := `中新网1月19日电 据韩媒报道，18日下午，韩国“我们共和党”在釜山举行集会，大批群众参与游行，高喊口号支持韩国前总统朴槿惠。游行队伍行进途中，一车辆突然闯入，造成7人受伤。
　　当地时间18日下午4点20分左右，游行队伍在釜山东区行进途中，A某驾驶的轿车在右转时撞上了游行人员。事故造成7名参加集会的人员(3名男性，4名女性)受伤，被送往附近医院接受治疗。
　　一位目击者称，“在街头游行过程中，车辆突然闯进游行队伍，撞伤了参与集会的人”。
　　据悉，警方正在对司机A某进行调查，查明准确的事故原因。
　　据报道，当地时间18日中午12时30分，由亲朴人士组成的韩国“我们共和党”主办的第167次“太极旗集会”在釜山站广场举行。参与者们在结束集会后，沿着车道游行，高喊口号“文在寅下台”、“朴槿惠无罪”、“弹劾无效”等。`
	ps := analyzer.Analysis(docid, doc)
	t.Log(ps)
	term := "人"
	length := len(term)
	t.Log(strings.Count(doc, term))
	for _, posting := range ps {
		if posting.Term == term {
			t.Log(strings.Count(doc, term) == len(posting.Offset))
			for _, offset := range posting.Offset {
				t.Log(doc[offset : offset+length])
			}
		}
	}
}*/

func TestAnalysis(t *testing.T) {
	t.Log("test2")
	stopWordPath := "/Users/wengguan/search_code/search/search_engine/configs/stop_word.txt"
	analyzer := AnalyzerFactory(stopWordPath)
	var docid uint64 = 0
	doc := objs.Doc{Ident: "88.199.1/aaa.def", Data: objs.Data{Modified: "北京市丰台区", Saled: "北京市海淀区", CreatedAt: time.Now()}}
	ps := analyzer.Analysis(docid, doc)
	t.Log(ps)

	docid++
	doc = objs.Doc{Ident: "88.199.1/bbb.def", Data: objs.Data{Modified: "上海市黄浦区", Saled: "上海市浦东新区", CreatedAt: time.Now()}}
	ps = analyzer.Analysis(docid, doc)
	t.Log(ps)

	docid++
	doc = objs.Doc{Ident: "88.199.1/ccc.def", Data: objs.Data{Modified: "天津市河东区", Saled: "天津市河西区", CreatedAt: time.Now()}}
	ps = analyzer.Analysis(docid, doc)
	t.Log(ps)

	docid++
	doc = objs.Doc{Ident: "88.199.1/ddd.def", Data: objs.Data{Modified: "江西省南昌市", Saled: "江西省吉安市", CreatedAt: time.Now()}}
	ps = analyzer.Analysis(docid, doc)
	t.Log(ps)

	docid++
	doc = objs.Doc{Ident: "88.199.1/eee.def", Data: objs.Data{Modified: "江西省南昌市", Saled: "江西省井冈山市", CreatedAt: time.Now()}}
	ps = analyzer.Analysis(docid, doc)
	t.Log(ps)
}
