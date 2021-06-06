package logparser

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"regexp"
)

//根据分组名 把键值映射到Map中
func MapGroup(list []string, names []string) map[string]interface{} {
	ret := make(map[string]interface{})

	if len(list) <= 1 {
		return nil
	}
	for index, name := range names {
		if index == 0 || name == "" {
			continue
		}
		ret[name] = list[index]
	}
	return ret
}

type HttpdParser struct {
	lines []string
}

func NewHttpdParser(filename string) *HttpdParser {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("logfile error:", err)
	}
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &HttpdParser{lines: lines}
}
func (this *HttpdParser) Parse() []map[string]interface{} {
	var buffer bytes.Buffer
	buffer.WriteString(`^(?P<ip>\d+.\d+.\d+.\d+).*?`)              // ip 格式
	buffer.WriteString(`\[(?P<time>.+?)\]\s*`)                     //time
	buffer.WriteString(`\"(?P<method>[A-Z]+)\s*(?P<url>.*?)\"\s*`) //method和URL
	buffer.WriteString(`(?P<status>\d+)\s*`)                       //status
	buffer.WriteString(`(?P<duration>\d+)\s*`)                     //duration
	ret := make([]map[string]interface{}, 0)
	for _, line := range this.lines {
		reg := regexp.MustCompile(buffer.String())
		result := reg.FindStringSubmatch(line)
		m := MapGroup(result, reg.SubexpNames())
		ret = append(ret, m)
	}
	return ret
}
