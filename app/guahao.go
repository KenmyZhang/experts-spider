package app

import(
	l4g "github.com/alecthomas/log4go"
	"regexp"	
	"strconv"
	"runtime"
)

var catos = []string{ "甲亢","乙肝","胃炎","高血脂","脑梗塞","癫痫","强直性脊柱炎","痛风","哮喘","肺结核", "胃溃疡", "克罗恩病", "肝硬化", "脑梗塞", "癫痫", "面瘫", "三叉神经痛", "老年痴呆", "脑血栓形成", "帕金森病", "脑供血不足", "失眠", "面肌抽搐", "重症肌无力", "高尿酸血症", "痛风", "甲状腺炎", "糖尿病足", "肥胖症", "低血糖", "甲减", "强直性脊柱炎", "风湿病", "类风湿性关节炎", "干燥综合征", "特发性关节痛综合征", "硬皮病", "系统性红斑狼疮", "白塞病", "外科常见病", "前列腺炎", "肺癌", "痔疮", "甲状腺腺瘤", "直肠癌", "肾结石", "结肠癌", "胆囊结石", "胃癌", "食管癌", "胆囊癌", "肝血管瘤", "早泄", "肾结石", "包皮过长", "性功能障碍", "男性不育症", "尿路感染", "前列腺增生", "前列腺癌", "肾积水", "乳腺增生", "乳腺纤维腺瘤", "乳腺腺病", "急性乳腺炎", "多乳房", "乳腺导管扩张症", "胃癌", "甲状腺腺瘤", "胰腺炎", "甲状腺癌", "胰腺癌", "脂肪瘤", "甲状腺结节样病变", "阑尾炎", "胰腺囊肿", "脾大", "肝肿瘤", "胆囊结石", "胆囊息肉", "肝血管瘤", "胆囊炎", "胆管结石", "胆管癌", "肝腹水", "胆囊癌", "胆总管结石", "宫颈癌", "月经不调", "输卵管堵塞", "阴道炎", "多囊卵巢综合征", "围绝经期综合征", "宫颈炎", "盆腔炎", "子宫内膜异位症", "子宫内膜增生", "宫外孕", "子宫腺肌症", "产后出血", "早产", "自然流产", "巧克力囊肿", "子宫脱垂", "宫颈癌前病变", "子宫颈息肉", "畸胎瘤", "卵巢癌", "子宫内膜息肉", "子宫内膜癌", "尖锐湿疣", "腋臭", "灰指甲", "胎记", "扁平疣", "带状疱疹", "甲沟炎", "黑色素瘤", "脂溢性皮炎", "梅毒", "寻常痤疮", "毛囊炎", "皮肤瘙痒症", "神经性皮炎", "外阴黏膜白斑", "表皮痣", "皮炎", "色素痣", "生殖器疱疹", "鱼鳞病", "脚气", "寻常疣", "骨质增生", "肩周炎", "腱鞘炎", "半月板损伤", "滑膜炎", "关节炎", "慢性腰背痛", "股骨头缺血性坏死", "骨折", "颈椎间盘突出症", "慢性腰背痛", "颈椎间盘突出症", "腰椎椎管狭窄症", "脊髓型颈椎病", "神经根型颈椎病", "脊柱骨折", "脊柱肿瘤", "肩周炎", "半月板损伤", "滑膜炎", "痛风性关节炎", "肩袖损伤", "肱骨外上髁炎", "骨关节炎", "先天性髋关节脱位", "肌腱炎", "骨折", "跟腱炎", "半月板撕裂", "跟腱断裂", "腓总神经损伤", "髌骨骨折", "股骨颈骨折", "踝关节骨折", "膝外侧韧带损伤", "鼻炎", "鼻窦炎", "中耳炎", "三叉神经痛", "鼻咽癌", "鼻中隔偏曲", "扁桃体炎", "鼻息肉", "声带息肉", "美尼尔综合症", "急性鼻炎", "喉癌", "鼓膜穿孔", "咽异感症", "慢性咽炎", "白内障", "青光眼", "近视", "弱视", "结膜炎", "黄斑变性", "玻璃体混浊", "斜视", "角膜炎", "视网膜脱落", "散光", "葡萄膜炎", "睑腺炎", "沙眼", "视神经炎", "甲状腺癌", "腮腺炎", "咽部肿瘤", "颈动脉瘤", "心肌炎", "心律失常", "心肌缺血", "房颤", "心脏神经症", "扩张型心肌病", "肺心病", "风湿性心脏病", "小儿哮喘", "小儿癫痫", "小儿抽动症", "小儿支气管肺炎", "小儿原发性肾病综合征", "唐氏综合征", "先天性斜颈", "牙周炎", "龋齿", "牙龈炎", "口腔扁平苔藓", "牙齿磨损", "牙髓炎", "颞下颌关节紊乱综合征", "口腔癌", "口腔黏膜白斑", "抑郁症", "小儿多动症", "焦虑症", "强迫性障碍", "精神分裂症", "心理障碍", "产后抑郁症", "神经衰弱"}
var provinces = []string{"上海&pi=2", "北京&pi=1", "广东&=29", "江苏&=22", "浙江&=24", "陕西&=9", "甘肃&=11", "山东&=21", "山西&pi=8", "湖北&pi=19", "湖南&pi=30", "天津&pi=3", "四川&pi=15", "江西&pi=25", "安徽&pi=23", "河南&pi=20", "河北&pi=16", "青海&pi=12", "辽宁&pi=5","贵州&pi=18", "重庆&pi=4", "黑龙江&pi=7", "云南&pi=17", "广西&pi=31", "宁夏&pi=10", "西藏&pi=14", "内蒙古&pi=33", "海南&pi=32", "吉林&pi=6", "新疆&pi=13", "福建&pi=27"}

var(
	doctorUrlReg       = regexp.MustCompile(`<a target="_blank" monitor="search_allpg,search_allpg,doctor" monitor-doctor-id="[\S]+?" href="[\S]+?"`)
	doctorUrlRegPrefix = regexp.MustCompile(`<a target="_blank" monitor="search_allpg,search_allpg,doctor" monitor-doctor-id="[\S]+?" href="`)
	doctorUrlRegSuffix = regexp.MustCompile(`"`)

	totalOfSearchResultReg       = regexp.MustCompile(`<span class="result-num">找到<strong id="J_ResultNum">[\s]+[0-9]+?[\s]+</strong>位医生</span>`)
	totalOfSearchResultRegPrefix = regexp.MustCompile(`<span class="result-num">找到<strong id="J_ResultNum">[\s]+`)
	totalOfSearchResultRegSuffix = regexp.MustCompile(`[\s]+</strong>位医生</span>`)	


	hospitalReg = regexp.MustCompile(`<div class="item hospital">[\s]+<label for="">医院：</label>[\s\S]+?</div>[\s]+</div>`)
	hospitalAllReg = regexp.MustCompile(`title="[\s\S]+?"`)
	hospitalAllRegPrefix = regexp.MustCompile(`title="`)
	hospitalAllRegSuffix = regexp.MustCompile(`"`)


	deptReg = regexp.MustCompile(`<div class="item dept">[\s]+<label for="">科室：</label>[\s]+<div class="contain" id="schedules-dept">[\s\S]+?</div>[\s]+</div>`)
	deptRegAllReg = regexp.MustCompile(`<a href="javascript:;" data-deptId="[\s\S]+?" data-hospId="[\s\S]+?>[\s]+[\S]+[\s]+</a>`)
	deptRegAllRegPrefix = regexp.MustCompile(`<a href="javascript:;" data-deptId="[\s\S]+?" data-hospId="[\s\S]+?>[\s]+`)
	deptRegAllRegSuffix = regexp.MustCompile(`[\s]+</a>`)	


	jobTitleReg = regexp.MustCompile(`<strong class="J_ExpertName">[\s\S]+?</strong>[\s\S]+?<a href=`)
	jobTitleRegPrefix = regexp.MustCompile(`<strong class="J_ExpertName">[\s\S]+?</strong>`)
	jobTitleRegSuffix = regexp.MustCompile(`<a href=`)
	nullReg = regexp.MustCompile(`\s`)
	spanReg = regexp.MustCompile(`<span>`)
	spanReg2 = regexp.MustCompile(`</span>`)

	doctorNameReg0 = regexp.MustCompile(`<strong class="J_ExpertName">[\s\S]+?</strong>`)
	doctorNameReg0Prefix = regexp.MustCompile(`<strong class="J_ExpertName">`)
	doctorNameReg0Suffix = regexp.MustCompile(`</strong>`)

	typeJudgeReg = regexp.MustCompile(`<b>擅长：</b>[\s]+<span>[\s\S]+?</span>[\s]+?</div>[\s]+?<div class="about">`)

	goodAtReg = regexp.MustCompile(`<b>擅长：</b>[\s]+<span>[\s\S]+?</span>[\s]+<a href="[\s\S]+?" data-description="[\s\S]+?">`)
	goodAtRegPrefix = regexp.MustCompile(`<b>擅长：</b>[\s]+<span>[\s\S]+?</span>[\s]+<a href="[\s\S]+?" data-description="`)
	goodAtRegSuffix = regexp.MustCompile(`">`)

	goodAtReg2 = regexp.MustCompile(`<b>擅长：</b>[\s]+<span>[\s\S]+?</span>`)
	goodAtReg2Prefix = regexp.MustCompile(`<b>擅长：</b>[\s]+<span>`)
	goodAtReg2Suffix = regexp.MustCompile(`</span>`)

	briefReg0 = regexp.MustCompile(`<b>简介：</b>[\s]+<span>[\s\S]+?</span>[\s]+<a href="[\s\S]+" data-description="[\s\S]+?" monitor="doctor,info,introduce_more">`)
	briefReg0Prefix = regexp.MustCompile(`<b>简介：</b>[\s]+<span>[\s\S]+?</span>[\s]+<a href="[\s\S]+" data-description="`)
	briefReg0Suffix = regexp.MustCompile(`" monitor="doctor,info,introduce_more">`)

	briefReg2 = regexp.MustCompile(`<b>简介：</b>[\s]+<span>[\s\S]+?</span>`)
	briefReg2Prefix = regexp.MustCompile(`<b>简介：</b>[\s]+<span>`)
	briefReg2Suffix = regexp.MustCompile(`</span>`)
)

func GetExpertFromGuaHao(urlChan chan string ) {
	for _, cato := range catos {
		for _, province := range provinces {
			url := `https://www.guahao.com/search/expert?iSq=&fhc=&fg=&q=` + cato + `&p=` + province + `&ci=all&c=不限&o=all&es=all&hl=all&ht=all&hk=&dt=&dty=&hdi=&mf=true&fg=0&ipIsShanghai=false&searchAll=Y&hospitalId=&standardDepartmentId=&consult=&volunteerDoctor=&imagetext=&phone=&diagnosis=&sort=general&hydate=all&activityId=&weightActivity=&weightWelifeTopic=&standardDepartmentId=&isStaticUrl=&staticUrl=&p1=p1&isDiseaseUrl=`
			//l4g.Debug(url + "-----begin")
			respBody, err := httpGet(url, false)
			if err != nil {
				l4g.Error("error:", url)
				continue
			}
			//l4g.Debug(url + "-----end")

			total := totalOfSearchResultReg.FindString(respBody)
    		total = totalOfSearchResultRegPrefix.ReplaceAllString(total, "")
    		total = totalOfSearchResultRegSuffix.ReplaceAllString(total, "")
    		count, _ := strconv.Atoi(total)
    		l4g.Debug(url + ", " + total)
    		l4g.Debug(url + "======(count/16 +1)=====" + strconv.Itoa(count/16 +1))
    		for i := 1; i <= (count/16 +1); i ++ {
    			inputUrl := url + "&pageNo=" + strconv.Itoa(i)
    			l4g.Debug(url)
    			GetDoctorUrlsFromPerPage(inputUrl, urlChan)
    		}
    	}
    }
    l4g.Error("Finish range cato")
}

func GetDoctorUrlsFromPerPage(inputUrl string, urlChan chan string) {
	runtime.Gosched()
	respBody, err := httpGet(inputUrl, false)
	if err != nil {
		l4g.Error(err.Error())
		return
	}

	doctorUrls := doctorUrlReg.FindAllString(respBody, -1)
    for _, url := range doctorUrls {
    	url = doctorUrlRegPrefix.ReplaceAllString(url, "")
    	url = doctorUrlRegSuffix.ReplaceAllString(url, "")
    	l4g.Debug(url)
    	urlChan <- url
    }
}

func GetDoctorDetails(urlChan chan string) {
	for {
		select {
		case url, ok := <-urlChan:
			if !ok {
				l4g.Error("Finish All")
				return
			} 
			//l4g.Debug(url + "-----begin")
			respBody, err := httpGet(url, false)		
			if err != nil {
				l4g.Error(err.Error())
				break
			}	
			//l4g.Debug(url + "------end")

			expertDetailData := &ExpertDetailData{}

			name := doctorNameReg0.FindString(respBody)
			name = doctorNameReg0Prefix.ReplaceAllString(name, "")
			name = doctorNameReg0Suffix.ReplaceAllString(name, "")
			expertDetailData.Name = name

			hospital := hospitalReg.FindString(respBody)
			hospitalAll := hospitalAllReg.FindAllString(hospital, -1)
			hospitalMap := map[string] string{}
			var hospitalName string
			for _, item :=  range hospitalAll {
				item = hospitalAllRegPrefix.ReplaceAllString(item, "")				
				item = hospitalAllRegSuffix.ReplaceAllString(item, "")
				if _, ok := hospitalMap[item]; !ok {
					hospitalMap[item] = ""
					hospitalName = hospitalName + item + " "
				}
			}
			expertDetailData.HospitalName = hospitalName

			dept := deptReg.FindString(respBody)
			deptAll := deptRegAllReg.FindAllString(dept, -1)
			deptMap := map[string]string{}
			var deptName string
			for _, item :=  range deptAll {
				item = deptRegAllRegPrefix.ReplaceAllString(item, "")				
				item = deptRegAllRegSuffix.ReplaceAllString(item, "")
				if _, ok := deptMap[item]; !ok {
					deptMap[item] = ""
					deptName = deptName + item + " "
				}
			}						
			expertDetailData.MedicalBranchName = deptName


			jobTitle := jobTitleReg.FindString(respBody)
			jobTitle = jobTitleRegPrefix.ReplaceAllString(jobTitle, "")
			jobTitle = jobTitleRegSuffix.ReplaceAllString(jobTitle, "")
			jobTitle = nullReg.ReplaceAllString(jobTitle, "")
			jobTitle = spanReg.ReplaceAllString(jobTitle, "")
			jobTitle = spanReg2.ReplaceAllString(jobTitle, "")
			expertDetailData.JobTitle = jobTitle

			var goodAt string
			if "" == typeJudgeReg.FindString(respBody) { 
				goodAt = goodAtReg.FindString(respBody)
				goodAt = goodAtRegPrefix.ReplaceAllString(goodAt, "")
				goodAt = goodAtRegSuffix.ReplaceAllString(goodAt, "")
			} else {
				goodAt = goodAtReg2.FindString(respBody)
				goodAt = goodAtReg2Prefix.ReplaceAllString(goodAt, "")
				goodAt = goodAtReg2Suffix.ReplaceAllString(goodAt, "")
			}
			expertDetailData.Skill = goodAt


			brief := briefReg0.FindString(respBody)
			brief = briefReg0Prefix.ReplaceAllString(brief, "")
			brief = briefReg0Suffix.ReplaceAllString(brief, "")

			if brief == "" {
				brief = briefReg2.FindString(respBody)
				brief = briefReg2Prefix.ReplaceAllString(brief, "")
				brief = briefReg2Suffix.ReplaceAllString(brief, "")
			}
			expertDetailData.Brief = brief
			if expertDetailData.Name != "" {
				SaveExpertDetail2(expertDetailData)
			}
		}
	}
}