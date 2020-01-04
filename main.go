package main

import (
	"log"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", simulateCar)
	log.Println(http.ListenAndServe(":8000", nil))
}

func simulateCar(res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("<html>"))

	// make a new car
	car := NewCar(res)

	res.Write([]byte("Bruce Wayne goes to the dealership and buys himself a 1998 " + car.Color + " " + car.Company + " " + car.Model + "<br>\n"))

	//begin car simulation
	for {

		//Park the car in the garage after 30 years
		car.Age()
		if car.Years >= 30 {
			break
		}

		//Change tires from drifting
		if car.Tires > 0 {
			res.Write([]byte("Bruce has to change the tires from being a boss<br>\n"))
			car.Tires = 0
			continue
		}

		//Change the oil when car hits 3000 miles
		if car.Mileage >= 3000 {
			res.Write([]byte("Time for Bruce to change the oil<br>\n"))
			car.Oil()
			continue
		}

		choiceNumber := rand.Intn(4) // 4 being the number of things that bruce does with the car
		switch choiceNumber + 1 {

		case 1:
			car.Drive()

		case 2:
			car.Park()

		case 3:
			car.Wash()

		case 4:
			car.Drift()
		}

	}
	res.Write([]byte("Bruce drove the crap out of that Skyline.....let's not run it in the gournd...time to park it...she is only for shows now<br>\n"))

}
