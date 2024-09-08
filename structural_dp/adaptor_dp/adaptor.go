package adaptor

import "fmt"

// target
type Mobile interface {
	ChargeAppleMobile()
}

// Concrete ProtoType implementation
type Apple struct{}

func (a *Apple) ChargeAppleMobile() {
	fmt.Println("Charging Apple Mobile!")
}

// Client
type Client struct{}

func (c *Client) ChargeMobile(mob Mobile) {
	mob.ChargeAppleMobile()
}

// Extended Implementation using Adapter

// Adaptee implementation
type Android struct{}

func (a *Android) ChargeAndroidMobile() {
	fmt.Println("Charging Android Mobile!")
}

// Adaptor Implementation

type Aaptor struct {
	Andorid *Android
}

func (ad *Aaptor) ChargeAppleMobile() {
	ad.Andorid.ChargeAndroidMobile()
}
