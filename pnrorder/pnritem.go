package pnrorder

import (
	"fmt"
	"strings"
)

type PnrItemObj interface {
	IsMatch(line string) bool
	Add(line string)
	Append(line string)
	Data() interface{}
	Name() string
}

type PNR struct {
	Text   string
	Code   string
	ObjAry []PnrItemObj
}

func NewPNR(code string) *PNR {

	ret := &PNR{
		Text: code,
	}
	ret.ObjAry = append(ret.ObjAry, NewPassengerItem(), NewFltSegmentItem(), NewOtherItem())
	return ret
}

func (p *PNR) Analysis() {
	p.Text = strings.TrimSpace(p.Text)
	if p.Text == "" {
		return
	}
	lines := strings.Split(p.Text, "\n")
	//第一行必须是--- TST开头
	if strings.HasPrefix(lines[0], "---") == false {
		return
	}
	//第二、三行为基本信息，从第二行取PNR编码
	l := strings.TrimSpace(lines[1])
	p.Code = l[len(l)-6:]

	fmt.Println(p.Code)

	eachItem(lines[3:], p.ObjAry)
}

func (p *PNR) Ouput() map[string]interface{} {
	rev := make(map[string]interface{})
	rev["code"] = p.Code
	for _, o := range p.ObjAry {
		n := o.Name()
		if n != "" {
			rev[n] = o.Data()
		}
	}
	return rev
}

func eachItem(strAry []string, objs []PnrItemObj) {

	var current PnrItemObj
	for _, row := range strAry {
		row = strings.TrimSpace(row)
		for _, o := range objs {
			if o.IsMatch(row) {
				current = o
				o.Add(row)
				row = ""
				break
			}
		}
		if row != "" {
			current.Append(row)
		}
	}
}
