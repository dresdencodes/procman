package procman

import (
	"log"
	"time"
	"strings"
)

func (m *manager) Opts() {
	if m.once() {return} 
	m.every()
}

func (m *manager) once() bool {
	if _, ok := m.findOpt("once"); ok {
		return true
	}
	return false
}

func (m *manager) findOpt(pre string) (string, bool) {
	tspl := strings.Split(m.t, " ")
	for key := range tspl { word := tspl[key]
		if strings.HasPrefix(word, "["+pre) {
			if pre == "every" {word = word + " " + tspl[key+1]}
			return strings.Trim(word, "[]"), true
		}
	}
	return "", false
}

func (m *manager) every() {

	// if select loop, no every
	if _,ok := m.findOpt("select"); ok {return}

	var everyFound bool
	var everyStr string
	var everySpl []string
	everyStr,everyFound = m.findOpt("every")
	
	// no sleep
	if !everyFound {
		m.fn = func() {
			for {
				m.fn()
			}
		}
		return
	}

	// define every spl
	everySpl = strings.Split(everyStr, " ")

	// make sure len is 2
	if len(everySpl) < 2 || len(everySpl) > 2 {
		log.Fatal("dgman_opts invalid opt length [every n] opt passed in title: ", m.t, " every split: ", everySpl)
	}

	// parse duration
	dur, err := time.ParseDuration(everySpl[1])
	if err!=nil {
		log.Fatal("dgman_opts error parsing [every n] sleep time in:", m.t)
	}
	
	// copy fn -> must be copied 
	fn := m.fnCopy()

	// with sleep
	m.fn = func(){
		for {
			fn()
			time.Sleep(dur)
		}
	}
	
}


func (m *manager) fnCopy() func() {
	fn := m.fn
	return fn
}