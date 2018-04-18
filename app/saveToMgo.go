package app

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	l4g "github.com/alecthomas/log4go"
)

var ExpertDetailsCollection *mgo.Collection
var ExpertDetailsCollection1 *mgo.Collection
var ExpertDetailsCollection2 *mgo.Collection
var MgoSession *mgo.Session

func init() {
	MgoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional Switch the MgoSession to a monotonic behavior.
	MgoSession.SetMode(mgo.Monotonic, true)

	ExpertDetailsCollection = MgoSession.DB("spider").C("expert_details0")
	ExpertDetailsCollection1 = MgoSession.DB("spider").C("expert_details1")
	ExpertDetailsCollection2 = MgoSession.DB("spider").C("expert_details2")
	ExpertDetailsCollection2.EnsureIndex(mgo.Index{
		Key:    []string{"url"},
		Unique: true,
	})

}

func SaveExpertDetail(expertDetailData *ExpertDetailData) {
	expertDetailData.PreSave()
	err := ExpertDetailsCollection.Insert(expertDetailData)
	if err != nil {
		l4g.Error(err)
		return
	}
}

func SaveExpertDetail1(expertDetailData *ExpertDetailData) {
	expertDetailData.PreSave()
	err := ExpertDetailsCollection1.Insert(expertDetailData)
	if err != nil {
		l4g.Error(err)
		return
	}
}

func SaveExpertDetail2(expertDetailData *ExpertDetailData) {
	expertDetailData.PreSave()
	err := ExpertDetailsCollection2.Insert(expertDetailData)
	if err != nil {
		l4g.Error(err)
		return
	}
}

