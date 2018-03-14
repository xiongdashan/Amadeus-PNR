package pnrorder

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type PassengerItem struct {
	Regex      *regexp.Regexp
	RegexDoc   *regexp.Regexp
	RegexTktNo *regexp.Regexp
	Text       string
	Docs       []string
	TktNos     []string
	Type       string
}

//乘客信息
type Passenger struct {
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Type           string `json:"type"`
	Gender         string `json:"gender"`
	IDCardType     string `json:"id_card_type"`
	IDCardNO       string `json:"id_card_no"`
	Birthday       string `json:"birthday"`
	Nationality    string `json:"nationality"`
	IDIssueCountry string `json:"id_issue_country"`
	IDExpireDate   string `json:"id_expireDate"`
	TicketNumber   string `json:"ticket_number"`
	Index          int    `json:"index"`
}

//1.ZHENG/HUABIN MR(INFZHHENG/BING MSTR/23NOV16)

func NewPassengerItem() *PassengerItem {
	return &PassengerItem{
		Regex:      regexp.MustCompile(`(\d+)\.[A-Z]+\/[A-Z]+(.*[^\d])$`),
		RegexDoc:   regexp.MustCompile(`(\d+)(\s)SSR(\s)DOCS(.+)`),
		RegexTktNo: regexp.MustCompile(`(\d)+\s+FA\sPAX\s(.+)`),
	}
}

func (p *PassengerItem) Name() string {
	return "passenger"
}

func (p *PassengerItem) IsMatch(line string) bool {
	if p.Regex.MatchString(line) {
		p.Type = "P"
		return true
	}
	if p.RegexDoc.MatchString(line) {
		p.Type = "D"
		return true
	}
	if p.RegexTktNo.MatchString(line) {
		p.Type = "T"
		return true
	}
	// p.Type = ""
	return false
}

func (p *PassengerItem) Add(line string) {
	if p.Type == "P" {
		p.Text = fmt.Sprintf("%s %s", p.Text, line)
		return
	}
	if p.Type == "D" {
		p.Docs = append(p.Docs, line)
		return
	}
	p.TktNos = append(p.TktNos, line)
}

func (p *PassengerItem) Append(line string) {
	if p.Type == "P" {
		p.Add(line)
	}
	if p.Type == "D" {
		pos := len(p.Docs) - 1
		p.Docs[pos] = p.Docs[pos] + " " + line
	}
	if p.Type == "T" {
		pos := len(p.TktNos) - 1
		p.TktNos[pos] = p.TktNos[pos] + " " + line
	}
}

func (p *PassengerItem) Data() interface{} {
	rev := splitNamesLine(p.Text, p.Docs, p.TktNos)
	return rev
}

// 分割人名，并从Docs项里获取证件信息
func splitNamesLine(line string, docs []string, tktNos []string) (rev []*Passenger) {
	//假设排序正确
	// TODO: 排标识准备排序
	line = strings.TrimSpace(line)
	ary := regexp.MustCompile(`\d\.`).Split(line, -1)

	// 成人
	regADT := regexp.MustCompile(`^([A-Z]+)/([A-Z\s]+)(\s(MS|MR))?$`)
	// 婴儿 ZHANG/HUABIN MR(INFZHHENG/BING MSTR/23NOV16)
	regINF := regexp.MustCompile(`^([A-Z]+)/([A-Z\s]+)(\s(MS|MR))?\(INF([A-Z]+)/([A-Z]+)\s(MSTR|MSSR)/(\d{2})([A-Z]{3})(\d{2})\)$`)
	// 儿童 ZHENG/HUA QING MSTR(CHD/20NOV10)
	regCHD := regexp.MustCompile(`([A-Z]+)/([A-Z\s]+)(\s(MSRT|MSST))?\(CHD/(\d{2})([A-Z]{3})(\d{2})\)`)
	for i, v := range ary {
		v = strings.TrimSpace(v)
		if len(v) == 0 {
			continue
		}
		fmt.Println(v)
		var p *Passenger

		//pos := i + 1

		//成人....
		if regADT.MatchString(v) {
			p = adtInfo(regADT, v, docs, i)
			// 票号
			p.splitTicketLine(tktNos)
			rev = append(rev, p)
			continue
		}

		//婴儿
		if regINF.MatchString(v) {
			pADT, pINF := infAndDdtInfo(regINF, v, i)
			p = pADT
			p.splitDoc(docs)
			rev = append(rev, p)

			pINF.splitDoc(docs)

			rev = append(rev, pINF)

			continue
		}

		// 儿童
		if regCHD.MatchString(v) {
			p = chdInfo(regCHD, v, i)
			p.splitDoc(docs)
			p.splitTicketLine(tktNos)
			rev = append(rev, p)
		}

		// rev = append(rev, p)
	}
	return
}

// ZHENG/HUA YOU MSRT(CHD/20NOV10)
func chdInfo(regCHD *regexp.Regexp, input string, pos int) (pCHD *Passenger) {
	match := regCHD.FindAllStringSubmatch(input, -1)
	m := match[0]
	lastName, gender := nameAndGender(m[2])
	birthday := fmt.Sprintf("%s%s%s", m[5], m[6], m[7])
	pCHD = &Passenger{}
	pCHD.Type = "CHD"
	pCHD.FirstName = m[1]
	pCHD.LastName = lastName
	pCHD.Gender = gender
	pCHD.Birthday = formatBirthday(birthday, formatDate)
	pCHD.Index = pos
	return
}

// 成人 + 婴儿 基本信息
func infAndDdtInfo(regINF *regexp.Regexp, input string, pos int) (pADT *Passenger, pINF *Passenger) {
	match := regINF.FindAllStringSubmatch(input, -1)
	fmt.Println(match)
	pADT = adtInfoFromMatch(match)
	pADT.Index = pos
	m := match[0]
	lastName, gender := nameAndGender(m[7])
	birthday := fmt.Sprintf("%s%s%s", m[8], m[9], m[10])
	pINF = &Passenger{}
	pINF.FirstName = m[5]
	pINF.LastName = lastName
	pINF.Type = "INF"
	pINF.Gender = gender
	pINF.Index = pos
	pINF.Birthday = formatBirthday(birthday, formatDate)
	return
}

// 性别
func nameAndGender(input string) (string, string) {
	input = strings.TrimSpace(input)
	dict := map[string]string{
		" MSTR": "M",
		" MSSR": "F",
		" MS":   "M",
		" MR":   "F",
	}

	for key, v := range dict {
		if strings.HasSuffix(input, key) {
			return strings.TrimSuffix(input, key), v
		}
	}

	return input, ""
}

// 成人匹配基本信息
func adtInfoFromMatch(match [][]string) (p *Passenger) {
	lastName, gender := nameAndGender(match[0][2])
	p = &Passenger{}
	p.FirstName = match[0][1]
	p.LastName = lastName
	p.Gender = gender
	return
}

// 找出当前成人的基本信息和证件信息
func adtInfo(regADT *regexp.Regexp, input string, docs []string, pos int) (p *Passenger) {
	match := regADT.FindAllStringSubmatch(input, -1)
	fmt.Println(match[0])
	//匹配基本信息
	p = adtInfoFromMatch(match)
	p.Index = pos
	p.Type = "ADT"
	//再分析Docs项
	p.splitDoc(docs)
	return
}

func (p *Passenger) splitTicketLine(tktNos []string) {
	if tktNos == nil || len(tktNos) == 0 {
		return
	}
	suffix := fmt.Sprintf("/P%d", p.Index)
	for _, v := range tktNos {
		v := strings.TrimSpace(v)
		if strings.HasSuffix(v, suffix) {
			aryData := strings.Split(v, "/")
			reg := regexp.MustCompile(`(\d+)\s+FA\s+PAX`)
			number := reg.ReplaceAllString(aryData[0], "")
			fmt.Println(number)
			p.TicketNumber = strings.TrimSpace(number)
			break
		}
	}
}

//分析ＳＳＲ　ＤＯＣ项
// 　20 SSR DOCS HU HK1 ////28APR04/M//28DEC22/MA/GERRYYUCHEN/P3
// 10 SSR DOCS(乘客信息) AA HK1 P（护照）/CHN（国籍1）/PE14073897（护照号2）/CHN（签发国3）/22MAY60（生日4）/M（性别5）/27NOV22（有效期6）/SONG（7姓）/

func (p *Passenger) splitDoc(docs []string) {
	if docs == nil || len(docs) == 0 {
		return
	}
	suffix := fmt.Sprintf("/P%d", p.Index)
	for _, v := range docs {
		v = strings.TrimSpace(v)
		if strings.HasSuffix(v, suffix) {
			aryData := strings.Split(v, "/")
			genter := aryData[5]
			//如果带有婴儿标识，但乘客并非婴儿类型
			if strings.HasSuffix(genter, "I") && p.Type != "INF" {
				continue
			}
			//Amadeus证件类型只有护照
			p.IDCardType = "P"
			p.Nationality = aryData[1]
			p.IDCardNO = aryData[2]
			p.IDIssueCountry = aryData[3]
			p.Birthday = formatBirthday(aryData[4], formatDate)
			if len(aryData[5]) > 0 {
				p.Gender = aryData[5][0:1]
			}
			if len(aryData[6]) > 0 {
				p.IDExpireDate = aryData[6]
			}
		}
	}
}

func formatDate(timeVal string) time.Time {
	if len(timeVal) == 0 {
		return time.Time{}
	}
	format := "02Jan06"
	t, _ := time.Parse(format, timeVal)
	return t
}

//  CHD - 儿童（2 到11 岁）; INF - 婴儿（不超过2 岁）
func formatBirthday(input string, fn func(timeVal string) time.Time) string {
	t := fn(input)
	if (t.Year() - time.Now().Year()) > 0 {
		t = t.AddDate(-100, 0, 0)
	}
	if t.Year() == 1 {
		return ""
	}
	return t.Format("2006-01-02")
}
