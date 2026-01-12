package snippets

import (
	"fmt"
	"sort"
)

// Flight - a struct that
// contains information about flights
type Flight struct {
	Origin      string
	Destination string
	Price       int
}

type ByPrice []Flight

func (array ByPrice) Len() int {
	return len(array)
}

func (array ByPrice) Less(i, j int) bool {
	return array[i].Price < array[j].Price
}

func (array ByPrice) Swap(i, j int) {
	array[i], array[j] = array[j], array[i]

}

// SortByPrice sorts flights from highest to lowest
func SortByPrice(flights []Flight) []Flight {
	sort.Sort(ByPrice(flights))
	return flights
}

func printFlights(flights []Flight) {
	for _, flight := range flights {
		fmt.Printf("Origin: %s, Destination: %s, Price: %d \n", flight.Origin, flight.Destination, flight.Price)
	}
}

func LaunchSortByPrice() {
	flights := []Flight{
		{Price: 30},
		{Price: 20},
		{Price: 50},
		{Price: 1000},
	}

	printFlights(flights)

	sort.Sort(ByPrice(flights))

	printFlights(flights)
}
