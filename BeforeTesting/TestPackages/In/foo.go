package foo

import "log"

func myUnexportedFunc() {
	log.Println("I am a private func")
}
