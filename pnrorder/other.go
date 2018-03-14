package pnrorder

import (
	"regexp"
)

type OtherItem struct {
	Regex *regexp.Regexp
}

//只要是数字加空格开头的
func NewOtherItem() *OtherItem {
	return &OtherItem{
		Regex: regexp.MustCompile(`^(\d+)(\s)(.+)`),
	}
}

func (o OtherItem) Name() string {
	return ""
}

func (o OtherItem) IsMatch(line string) bool {
	return o.Regex.MatchString(line)
}

func (o OtherItem) Add(line string) {
}

func (o OtherItem) Append(line string) {

}

func (o OtherItem) Data() interface{} {
	return nil
}
