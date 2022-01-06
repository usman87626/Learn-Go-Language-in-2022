package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of Golang")

	//PRESENT TIME - yyyy-mm-dd 18:31:20.8887155 +0500 PKT m=+0.003991401
	presentTime := time.Now()
	//fmt.Println(presentTime)

	/*
		===============TIME FORMATTING in GOLANG=======================
		Time formatting in Golang is little weird or you can say there are some values that you have to remember to format time/date.
		For example:
		 - 01-02-2006 is always used when we want to format date in this format
		 - To format time, 03:04:05AM/03:04:05PM/15:04:05 is used
		 - To Format Day, Monday is used

		 You can read more about it here: https://pkg.go.dev/time
	*/
	fmt.Println(presentTime.Format("01-02-2006 03:04:05PM Monday "))

	//We can manually create some date as
	createDate := time.Date(2020, time.August, 01, 03, 03, 03, 2, time.UTC)
	//Formatting method is same as explained above.
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))
}
