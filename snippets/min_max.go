package snippets

import (
	"errors"
	"fmt"
)

type flight struct {
	Origin      string
	Destination string
	Price       int
}

func GetMinMax(flights []flight) (int, int, error) {

	if len(flights) == 0 {
		return 0, 0, errors.New("flights is empty")
	}

	min := flights[0].Price
	max := flights[0].Price

	for _, flight := range flights {

		if max < flight.Price {
			max = flight.Price
		}

		if min > flight.Price {
			min = flight.Price
		}
	}
	return min, max, nil
}

func LaunchFindMinMax() {
	flights := []flight{
		{Origin: "ROM", Destination: "NYC", Price: 500},
		{Origin: "MIL", Destination: "LON", Price: 200},
		{Origin: "PAR", Destination: "BER", Price: 150},
		{Origin: "ROM", Destination: "PAR", Price: 300},
	}

	min, max, err := GetMinMax(flights)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("Min price:", min)
	fmt.Println("Max price:", max)

	empty := []flight{}
	_, _, err = GetMinMax(empty)
	if err != nil {
		fmt.Println("on empty slice:", err)
	}
}
