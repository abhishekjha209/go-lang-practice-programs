package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("15:04:05 02-01-2006 Monday"))

	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}

// Formats used in time moudle:

// Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
// Basic short date "2006/01/02" gives "2015/02/25"
// AM/PM            "3PM==3pm==15h" gives "11AM==11am==11h"
// No fraction      "Mon Jan _2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
// 0s for fraction  "15:04:05.00000" gives "11:06:39.12340"
// 9s for fraction  "15:04:05.99999999" gives "11:06:39.1234"
