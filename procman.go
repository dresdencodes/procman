package procman

import (
	"log"
	"strings"
)

type manager struct {
	t  	 string
	opts map[string]interface{}
	fn func()
}

func Run(t string, fn func()) {

	// report
	log.Println("manager - "+t+" running")

	// build manager
	m := &manager{
		t:t,
		fn:fn,
	}

	// parse opts & run
	m.Opts()
	m.fn()
	
}

type InitFunc struct {
	Title 		string 		
	Fn 			func()
}
var initFuncs []InitFunc

func InitManagers() {
	for key := range initFuncs { 
		initF := initFuncs[key]
		Run(initF.Title, initF.Fn)
	}
}

func InitRun(t string, fn func()) struct{} {
	if strings.Contains(t, "[blocking]") {log.Fatal("dgman InitRun() cannot contain [blocking] option")}
	initFuncs = append(initFuncs, InitFunc{
		Title:t,
		Fn:fn,
	})
	return struct{}{}
}
