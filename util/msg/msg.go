package msg

import "fmt"

var HideInfoStatus bool

func HideInfo(status bool) {
	HideInfoStatus = status
}

func Info(msg ...interface{}) {
	if !HideInfoStatus {
		fmt.Println(msg)
	}
}
