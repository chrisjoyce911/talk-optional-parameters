package main

import "github.com/chrisjoyce911/talk-optional-parameters/strawman"

func main() {
	straw := strawman.New("AUD")
	straw = strawman.New("AUD", strawman.Quantity(30))
	straw = strawman.New("AUD", strawman.Quantity(30), strawman.UsedFor("barskets"))

	// Using a preset
	straw = strawman.New("AUD", strawman.ForFeed())

	straw.BuyUnits(10)

	straw.CurrentUnits()
	straw.CurrentValue()

	straw.MakeHay()

	straw.CurrentUnits()
	straw.CurrentValue()

}
