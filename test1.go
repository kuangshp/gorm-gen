package main

import (
	"fmt"
	"time"
)

func TimeToUnix(e time.Time) int64 {
	timeUnix, _ := time.Parse("2006-01-02 15:04:05", e.Format("2006-01-02 15:04:05"))
	return timeUnix.UnixNano() / 1e6
}
func main() {
	fmt.Println(TimeToUnix(time.Now()))
}
