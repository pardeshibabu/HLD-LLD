package main

import (
	"fmt"
	adaptor "hld_lld/structural_dp/adaptor_dp"
)

func main() {
	fmt.Println("Hello Low Level Designing World!")
	a := &adaptor.Apple{}
	c := &adaptor.Client{}
	c.ChargeMobile(a)

	// Extended Implementation
	androidObj := &adaptor.Android{}
	ad := &adaptor.Aaptor{
		Andorid: androidObj,
	}
	androidObj.ChargeAndroidMobile()
	c.ChargeMobile(ad)
}
