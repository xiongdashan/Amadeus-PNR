package pnrorder

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type FlightSegment struct {
	Index             int    `json:"index"`
	MarketingAirline  string `json:"marketing_airline"`
	FlightNumber      string `json:"flight_number"`
	ClassAvail        string `json:"class_avail"`
	DepartureDate     string `json:"departure_date"`
	ArrivalDate       string `json:"arrival_date"`
	Weekday           string `json:"weekday"`
	DepartureCityCode string `json:"departure_city_code"`
	DepartureCity     string `json:"departure_city"`
	ArrivalCityCode   string `json:"arrival_city_code"`
	ArrivalCity       string `json:"arrival_city"`
	DepartureTime     string `json:"departure_time"`
	ArrivalTime       string `json:"arrival_time"`
	BigCode           string `json:"big_code"`
}

type FlightSegmentItem struct {
	Regex    *regexp.Regexp
	Segments []string
}

//3  AF 381 L 23OCT 2*PEKCDG HK2  0100 0555  23OCT  E  AF/Q5735D
//匹配开头部分 3  AF 381 L 23
func NewFltSegmentItem() *FlightSegmentItem {
	return &FlightSegmentItem{
		Regex: regexp.MustCompile(`(\d+)(\s+)([A-Z]{2}|(\d\w))(\s+)?(\d+)(\s)?([A-Z]{1})\s+([0-9]{2})(.+)`),
	}
}

func (f *FlightSegmentItem) Name() string {
	return "flight_section"
}

func (f *FlightSegmentItem) IsMatch(line string) bool {
	return f.Regex.MatchString(line)
}

func (f *FlightSegmentItem) Add(line string) {
	f.Segments = append(f.Segments, line)
}

func (f *FlightSegmentItem) Append(line string) {
	pos := len(f.Segments) - 1
	f.Segments[pos] = f.Segments[pos] + " " + line
}

func (f *FlightSegmentItem) Data() interface{} {
	//return f.Segments
	data := pickOutSegments(f.Segments)
	return data
}

// (\d+)\s+([A-Z0-9\s]+)\s([A-Z]{1})\s((\d{2}[A-Z]{3}))\s(\d{1})(\*)?([A-Z]{6})\s([A-Z]{2}\d{1})\s+(\d{4})\s+(\d{4})\s+((\d{2}[A-Z]{3}))\s+E\s+([A-Z0-9]{2})\/([A-Z0-9]{6})
// 9 AF 381 L 23OCT 2*PEKCDG HK2  0100 0555  23OCT  E  AF/Q5735D
func pickOutSegments(segments []string) (ret []*FlightSegment) {
	regex := regexp.MustCompile(`(\d+)\s+([A-Z0-9\s]+)\s([A-Z]{1})\s((\d{2}[A-Z]{3}))\s(\d{1})(\*|\s+)?([A-Z]{6})\s([A-Z]{2}\d{1})\s+(\d{4})\s+(\d{4})\s+((\d{2}[A-Z]{3}))\s+E\s+([A-Z0-9]{2})\/([A-Z0-9]{6})`)

	for i, v := range segments {
		v = strings.TrimSpace(v)
		matches := regex.FindAllStringSubmatch(v, -1)
		if len(matches) == 0 {
			fmt.Println(v)
			continue
		}
		m := matches[0]
		fltInfo := m[2]

		fltSgm := &FlightSegment{}
		// TODO：按每行的开头数字重新排序
		fltSgm.Index = i + 1
		fltSgm.MarketingAirline = fltInfo[:2]
		fltSgm.FlightNumber = strings.TrimSpace(fltInfo[2:])
		fltSgm.ClassAvail = m[3]
		fltSgm.ArrivalDate = formatSegmentDate(m[4])
		fltSgm.Weekday = m[6]
		fltSgm.DepartureCityCode = m[8][:3]
		fltSgm.DepartureTime = formatSegmentTime(m[10])
		fltSgm.DepartureDate = formatSegmentDate(m[13])
		fltSgm.ArrivalCityCode = m[8][3:]
		fltSgm.ArrivalTime = formatSegmentTime(m[11])
		fltSgm.BigCode = m[15]

		ret = append(ret, fltSgm)
	}
	return

}

func formatSegmentTime(input string) string {
	return fmt.Sprintf("%s:%s", input[:2], input[2:])
}

// 航段日期 有可能存在跨年的情况
func formatSegmentDate(input string) string {
	format := "02Jan"
	t, _ := time.Parse(format, input)

	return fmt.Sprintf("%d-%02d-%02d", time.Now().Year(), t.Month(), t.Day())
}
