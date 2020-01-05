// https://www.codeabbey.com/index/task_view/modulo-and-time-difference--ru

package main

import (
	"fmt"
)

const (
	secInMinit = int64(60)
	secInHour  = secInMinit * 60
	secInDay   = secInHour * 24
)

func ScanInt64() (i_ int64) {
	fmt.Scanf("%d", &i_)
	return
}

func ScanTime() (dd_, hh_, mm_, ss_ int64) {
	dd_ = ScanInt64()
	hh_ = ScanInt64()
	mm_ = ScanInt64()
	ss_ = ScanInt64()
	return
}

func ComputeTimeStamp(dd, hh, mm, ss int64) (timeStamp_ int64) {
	timeStamp_ = ss + mm*secInMinit + hh*secInHour + dd*secInDay
	return
}

func TimeInterval(timeStampFrom, timeStampTo int64) (dd_, hh_, mm_, ss_ int64) {
	timeStampInterval := timeStampTo - timeStampFrom

	dd_ = timeStampInterval / secInDay
	timeStampInterval %= secInDay

	hh_ = timeStampInterval / secInHour
	timeStampInterval %= secInHour

	mm_ = timeStampInterval / secInMinit
	timeStampInterval %= secInMinit

	ss_ = timeStampInterval
	return
}

func main() {
	for i, count := int64(0), ScanInt64(); i < count; i++ {
		dd, hh, mm, ss := TimeInterval(ComputeTimeStamp(ScanTime()), ComputeTimeStamp(ScanTime()))
		fmt.Printf("(%v %v %v %v) ", dd, hh, mm, ss)
	}
}
