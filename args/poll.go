package args

import (
	"log"
	"os"
)

var Args []string

func init() {
	Args = os.Args[1:]
}

func basePoll() string {
	if len(Args) == 0 {
		log.Fatal("not enough args")
	}
	arg, newArgs := Args[0], Args[1:]
	Args = newArgs
	return arg
}

func Poll() string {
	return basePoll()
}

func PollBytes() []byte {
	return []byte(basePoll())
}
