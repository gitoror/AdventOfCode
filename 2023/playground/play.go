package main

import "fmt"

type Vehicule struct {
	name string
}

type Car struct {
	Vehicule
	car_type string
}

type Moto struct {
	Vehicule
	moto_type string
}

func main() {
	vehicules := []interface{}{}
	car := Car{}
	car.name = "car"
	car.car_type = "sport"
	vehicules = append(vehicules, car)
	vehicules = append(vehicules, Moto{Vehicule{"moto"}, "cross"})
	// for each vehicule in vehicules, if vehicule is a car, print car_type
	for _, vehicule := range vehicules {
		switch vehicule.(type) {
		case Car:
			fmt.Println(vehicule.(Car).car_type)
		case Moto:
			fmt.Println(vehicule.(Moto).moto_type)
		}
	}
}
