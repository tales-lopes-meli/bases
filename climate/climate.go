package climate

import "fmt"

func ShowClimate(temperature float32, moisture int, pressure float32){
	fmt.Printf("Temperature is %f, air moisture is %d and pressure is %f\n", temperature, moisture, pressure)
}