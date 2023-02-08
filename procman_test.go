package procman 

import (
	"log"
	"time"
    "testing"
)





// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test1(t *testing.T) {

    var _ = InitRun("[every 30s] test with delay", func() {
		log.Println("every 30s test ")
	})

	InitManagers()
	time.Sleep(200 * time.Second)
}
