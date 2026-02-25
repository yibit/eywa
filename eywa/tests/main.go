package main

import (
	"fmt"

	"github.com/6tail/tyme4go/tyme"
)

func main() {
	// 公历日
	solarDay, err := tyme.SolarDay{}.FromYmd(1986, 5, 29)

	// 参数有误无法创建公历日对象时，err错误信息不为空
	if err == nil {
		// 1986年5月29日
		fmt.Println(solarDay)

		// 农历丙寅年四月廿一
		fmt.Println(solarDay.GetLunarDay())

		// 第十七饶迥火虎年四月廿一
		day, _ := solarDay.GetRabByungDay()

		fmt.Println(day)
	}
}
