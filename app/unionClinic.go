package app 

import (
	l4g "github.com/alecthomas/log4go"
	"strconv"
	"strings"
	"regexp"	
)

var (
	doctorNameReg       = regexp.MustCompile(`<div class="doc-name">[\s\S]+?</div>`)
	doctorNameRegPrefix = regexp.MustCompile(`<div class="doc-name">`)
	doctorNameRegSuffix = regexp.MustCompile(`</div>`)

	medicalBranchNameReg       = regexp.MustCompile(`<div class="doc-dep">[\s]+<span>[\s\S]+?</span>[\s]+<span>[\s\S]+?</span>[\s]+</div>[\s]+<div class="doc-hospital">`)
	medicalBranchNameRegPrefix = regexp.MustCompile(`<div class="doc-dep">[\s]+<span>`)
	medicalBranchNameRegMiddle = regexp.MustCompile(`</span>[\s]+<span>`)
	medicalBranchNameRegSuffix = regexp.MustCompile(`</span>[\s]+</div>[\s]+<div class="doc-hospital">`)	

	hospitalNameReg       = regexp.MustCompile(`<div class="doc-hospital">[\s]+?所在医院：[\s\S]+?</div>`)
	hospitalNameRegPrefix = regexp.MustCompile(`<div class="doc-hospital">[\s]+?所在医院：`)
	hospitalNameRegSuffix = regexp.MustCompile(`</div>`)	
	space = regexp.MustCompile(` `)
	span = regexp.MustCompile(`<span>`)
	spanSlat = regexp.MustCompile(`</span>`)
	huanhang = regexp.MustCompile(`\n`)

	goodReg       = regexp.MustCompile(`<div class="doc-good">[\s]+<div class="item-title">[\s]+<img src="/assets/image/good_at.png" alt="">[\s]+擅长[\s]+</div>[\s]+<div class="item-desc">[\s\S]+?</div>[\s\S]+?</div>`)
	goodRegPrefix = regexp.MustCompile(`<div class="doc-good">[\s]+<div class="item-title">[\s]+<img src="/assets/image/good_at.png" alt="">[\s]+擅长[\s]+</div>[\s]+<div class="item-desc">`)
	goodRegSuffix = regexp.MustCompile(`</div>[\s\S]+?</div>`)

	briefReg       = regexp.MustCompile(`<div class="doc-intro">[\s]+<div class="item-title">[\s]+<img src="/assets/image/introduce.png" alt="">[\s]+专家介绍[\s]+</div>[\s]+<div class="item-desc">[\s\S]+?</div>`)
	briefRegPrefix = regexp.MustCompile(`<div class="doc-intro">[\s]+<div class="item-title">[\s]+<img src="/assets/image/introduce.png" alt="">[\s]+专家介绍[\s]+</div>[\s]+<div class="item-desc">`)
	briefRegSuffix = regexp.MustCompile(`</div>`)

		

)

func GetAllExpertId() {
	offset := 0
	for {
		if offset >5000 {
			l4g.Error("finish 0")
			return
		} 
		url := `https://www.unionclinic.cn/doctor/detail/` + strconv.Itoa(offset)
		respBody, err := httpGet(url, false)
		if err != nil {
			l4g.Error(err.Error())
			continue
		}
		offset++

		expertDetailData := &ExpertDetailData{}
		name := doctorNameReg.FindString(respBody)
		name = doctorNameRegPrefix.ReplaceAllString(name, "")
		name = doctorNameRegSuffix.ReplaceAllString(name, "")
		name = strings.TrimPrefix(name, " ")
		names := strings.Split(name, " ")
		if len(names) == 2 {
			expertDetailData.Name = names[0]
			expertDetailData.JobTitle = names[1]
		} else {
			expertDetailData.Name = names[0]
		}

		if expertDetailData.Name == "" {
			l4g.Error("name is null:", name)
			continue
		}

		medicalBranchName := medicalBranchNameReg.FindString(respBody)
		medicalBranchName = medicalBranchNameRegPrefix.ReplaceAllString(medicalBranchName, "")
		medicalBranchName = medicalBranchNameRegMiddle.ReplaceAllString(medicalBranchName, " ")
		medicalBranchName = medicalBranchNameRegSuffix.ReplaceAllString(medicalBranchName, "")
		expertDetailData.MedicalBranchName = medicalBranchName		

		hospitalName := hospitalNameReg.FindString(respBody)
		hospitalName = hospitalNameRegPrefix.ReplaceAllString(hospitalName, "")
		hospitalName = hospitalNameRegSuffix.ReplaceAllString(hospitalName, "")
		hospitalName = space.ReplaceAllString(hospitalName, "")
		hospitalName = span.ReplaceAllString(hospitalName, " ")
		hospitalName = spanSlat.ReplaceAllString(hospitalName, " ")
		hospitalName = huanhang.ReplaceAllString(hospitalName, " ")
		expertDetailData.HospitalName = hospitalName

		skill := goodReg.FindString(respBody)
		skill = goodRegPrefix.ReplaceAllString(skill, "")
		skill = goodRegSuffix.ReplaceAllString(skill, "")
		expertDetailData.Skill = skill 

		brief := briefReg.FindString(respBody)
		brief = briefRegPrefix.ReplaceAllString(brief, "")
		brief = briefRegSuffix.ReplaceAllString(brief, "")
		expertDetailData.Brief = brief
		SaveExpertDetail1(expertDetailData)
	}
}