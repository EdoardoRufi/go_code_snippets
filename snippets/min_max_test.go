package snippets

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func genFights(flightsNumber, maxFlightPrice int) []flight {

	var flights = make([]flight, flightsNumber)

	for i := 0; i < flightsNumber; i++ {
		flight := flight{"", "", rand.Intn(maxFlightPrice)}
		flights = append(flights, flight)
	}
	return flights
}

func BenchmarkGetMinMaxNormal(b *testing.B) {
	benchmarkGetMinMax(b, GetMinMax)
}

func benchmarkGetMinMax(b *testing.B, fn func([]flight) (int, int, error)) {
	var min int
	var max int
	flights := genFights(10000, 1000)
	fmt.Println(flights)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		min, max, _ = fn(flights)
	}
	println("min" + strconv.Itoa(min) + "max:" + strconv.Itoa(max))
}
