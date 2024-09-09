package main

import (
	"fmt"
	statedp "hld_lld/behaviour_dp/state_dp"
)

func main() {
	fmt.Println("Hello Low Level Designing World!")
	// ------------------------- Adaptor Design Pattern ----------------------------
	// a := &adaptor.Apple{}
	// c := &adaptor.Client{}
	// c.ChargeMobile(a)

	// // Extended Implementation
	// androidObj := &adaptor.Android{}
	// ad := &adaptor.Aaptor{
	// 	Andorid: androidObj,
	// }
	// androidObj.ChargeAndroidMobile()
	// c.ChargeMobile(ad)

	// ------------------------- State Design Pattern ----------------------------
	tvContext := statedp.GetContext() // Default state is Off
	tvContext.GetState()              // Get the State as Off

	tvOn := statedp.TvOn{}
	tvContext.SetState(&tvOn) // Change the current state as On
	tvContext.GetState()      // Get the Stat as On

}
