package main

import "fmt"

var (
	applePrice     float32 = 5.99
	pearPrice      float32 = 7
	myCash         float32 = 23
	amountOfApples int     = 9
	amountOfPears  int     = 8
	applesCount    int     = 2
	pearsCount     int     = 2
)

func main() {
	resultFirstQuestion := applePrice*float32(amountOfApples) + pearPrice*float32(amountOfPears)
	fmt.Println("1. Скільки грошей треба витратити, щоб купити", amountOfApples, "яблук та", amountOfPears, "груш?", "вiдповiдь:",
		resultFirstQuestion, "грн.")
	resultSecondQuestion := int(myCash / pearPrice)
	fmt.Println("2. Скільки груш ми можемо купити?", "вiдповiдь:", resultSecondQuestion, "шт.")
	resultThirdQuestion := int(myCash / applePrice)
	fmt.Println("3. Скільки яблук ми можемо купити?", "вiдповiдь:", resultThirdQuestion, "шт.")
	resultFourthQuestion := myCash >= float32(pearsCount)*pearPrice+float32(applesCount)*applePrice
	fmt.Println("4. Чи ми можемо купити", pearsCount, "груші та", applesCount, "яблука?", "вiдповiдь:", resultFourthQuestion)
}
