package gorm

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	SeqIndexRegexString = "\\[\\d+\\]"
)

// Copied from golint
var commonInitialisms = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UI", "UID", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
var commonInitialismsReplacer *strings.Replacer
var seqIndexRegexp *regexp.Regexp

func init() {
	var commonInitialismsForReplacer []string
	var err error = nil
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)

	seqIndexRegexp, err = regexp.Compile(SeqIndexRegexString)
	if err != nil {
		panic("init regexp failed!" + err.Error())
	}
}

type safeMap struct {
    m map[string]string
    l *sync.RWMutex
}

func (s *safeMap) Set(key string, value string) {
    s.l.Lock()
    defer s.l.Unlock()
    s.m[key] = value
}

func (s *safeMap) Get(key string) string {
    s.l.RLock()
    defer s.l.RUnlock()
    return s.m[key]
}

func newSafeMap() *safeMap {
    return &safeMap{l: new(sync.RWMutex), m: make(map[string]string)}
}


//var smap = map[string]string{}
var smap = newSafeMap()

func ToDBName(name string) string {
	//if v, ok := smap[name]; ok {
		//return v
	//}
	if v := smap.Get(name); v != "" {
		return v
	}

	value := commonInitialismsReplacer.Replace(name)
	buf := bytes.NewBufferString("")
	for i, v := range value {
		if i > 0 && v >= 'A' && v <= 'Z' {
			buf.WriteRune('_')
		}
		buf.WriteRune(v)
	}

	s := strings.ToLower(buf.String())
	//smap[name] = s
	smap.Set(name, s)
	return s
}

type expr struct {
	expr string
	args []interface{}
}

func Expr(expression string, args ...interface{}) *expr {
	return &expr{expr: expression, args: args}
}

func DBName(table_name string) (string, string) {
	sname := strings.Split(table_name, ".")
	if len(sname) == 1 {
		return "", sname[0]
	}
	return sname[0], sname[1]
}

func GetSeqInIndex(index_name string) (string, int, bool) {
	seqInIndex := seqIndexRegexp.FindString(index_name)
	if seqInIndex == "" {
		return index_name, 0, false
	}
	seqInIndex = strings.Split(strings.Split(seqInIndex, "]")[0], "[")[1]
	seqNo, _ := strconv.ParseInt(seqInIndex, 10, 64)
	realIndex := seqIndexRegexp.Split(index_name, 2)[0]
	return realIndex, int(seqNo), true
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
