package app

import(
	l4g "github.com/alecthomas/log4go"
	"strconv"
	"strings"
	"io"
	"encoding/json"
	"bytes"
	"encoding/base32"
	"github.com/pborman/uuid"
)

type ExpertIdResult struct {
	Data []ExpertData `json:"data"`
}

type ExpertData struct {
	UserId string `json:"userId"`
}

func ExpertFromJson(data io.Reader) *ExpertIdResult {
	decoder := json.NewDecoder(data)
	var o ExpertIdResult
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func GetExpertId(numStrChan chan string) {
	for i := 0; i < 300 ; i ++ {
		cato := strconv.Itoa(i)			
		offset := 0
loop:		
		for {
			url := `http://ssc.dazhuanjia.com/ssc/doctor/list/subject/` + cato  + `?category=HIM&offset=` + strconv.Itoa(offset) + `&limit=24`
			respBody, err := httpGet(url, false)
			if err != nil {
				l4g.Error(err.Error())
				continue
			}
			offset = offset + 24
			rst := ExpertFromJson(strings.NewReader(respBody))
			for _, data := range rst.Data {
				numStrChan <- data.UserId
			}

			if len(rst.Data) < 10 {
				break loop
			} 
		}
	}
	l4g.Error("close numStrChan")
	close(numStrChan)
	return	
}

type ExpertDetail struct {
	Data ExpertDetailData `json:"data"`
}

type ExpertDetailData struct {
	Id string `bson:"_id" json:"id"`
    Name string `bson:"name" json:"name"`
	HospitalName string `bson:"hospitalName" json:"hospitalName"`
	JobTitle string `bson:"jobTitle" json:"jobTitle"`
	MedicalBranchName string `bson:"medicalBranchName" json:"medicalBranchName"`
	Brief string `bson:"brief" json:"brief"`
	Skill string `bson:"skill" json:"skill"`
	URL string `bson:"url" json:"url"`
}

var encoding = base32.NewEncoding("ybndrfg8e234fdfsxot1uwisza345h769")

func NewId() string {
	var b bytes.Buffer
	encoder := base32.NewEncoder(encoding, &b)
	encoder.Write(uuid.NewRandom())
	encoder.Close()
	b.Truncate(26) // removes the '==' padding
	return b.String()
}


func (u *ExpertDetailData) PreSave() {
	if u.Id == "" {
		u.Id = NewId()
	}
}

func ExpertDetailFromJson(data io.Reader) *ExpertDetail {
	decoder := json.NewDecoder(data)
	var o ExpertDetail
	err := decoder.Decode(&o)
	if err == nil {
		return &o
	} else {
		return nil
	}
}

func GetExpertDetail(numStrChan chan string) {
	for {
		if val, ok := <-numStrChan; !ok {
			l4g.Error("finish 1")
			return
		} else {
			url := `http://rbac-new.dazhuanjia.com/bdc/doctor/` + val + `/detail`
			respBody, err := httpGet(url, false)
			if err != nil {
				l4g.Error("error:", url)
				continue
			}
			rst := ExpertDetailFromJson(strings.NewReader(respBody))			
			SaveExpertDetail(&rst.Data)
		}

	}
}