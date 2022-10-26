package main

import "fmt"

func main() {
	profile := MakeProfile("Sasuke")
	fmt.Println("Name : ", profile.Name)
	fmt.Println("Health : ", profile.Health)
	fmt.Println("Power :", profile.Power)
	fmt.Println("Exp :", profile.Exp)
	fmt.Println("=====HERO POWER UP====")
	PowerUp := PowerUp(profile, 3)
	fmt.Println("Name : ", PowerUp.Name)
	fmt.Println("Health : ", PowerUp.Health)
	fmt.Println("Power : ", PowerUp.Power)
	fmt.Println("Exp : ", PowerUp.Exp)

}

func MakeProfile(CharName string) Profile {
	var Char Profile
	Char.Name = CharName
	if CharName == "Sasuke" {
		Char.Health = 200
		Char.Power = 100
		Char.Exp = 0
	}
	if CharName == "Goku" {
		Char.Health = 400
		Char.Power = 300
		Char.Exp = 100
	}
	if CharName == "Naruto" {
		Char.Health = 150
		Char.Power = 200
		Char.Exp = 50
	}

	return Char
}

func PowerUp(profile Profile, multiplier int) Profile {
	var CharPowerUp Profile
	CharPowerUp.Name = profile.Name
	CharPowerUp.Health = profile.Health + profile.Health*multiplier
	CharPowerUp.Power = profile.Power + profile.Power*multiplier
	CharPowerUp.Exp = profile.Exp + profile.Exp*multiplier
	return CharPowerUp
}

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}
