package state

import (
	"log"
	"sync"
	"tech_test/util"
)

type Adjustment struct {
	Multiplier int64
	Divider    int64
	Addend     int64
}

type State struct {
	Mux         sync.Mutex
	HandleCount int64
	Details     map[string][]int64
	Adjustments map[string][]Adjustment
}

var instance *State
var once sync.Once

func GetInstance() *State {
	once.Do(func() {
		instance = &State{
			Details:     make(map[string][]int64),
			Adjustments: make(map[string][]Adjustment),
		}
	})
	return instance
}

const logSalesInterval = 1       //10
const logAdjustmentsInterval = 5 // 50

func DoHouseKeeping(theState *State) {
	theState.HandleCount++
	if (theState.HandleCount % logSalesInterval) == 0 {
		for k, v := range theState.Details {
			log.Printf("Product: %v number of sales: %v", k, len(v))
			for i, el := range v {
				log.Printf("  Sale: %v value: %v", i, el)
			}
		}
	}
	if (theState.HandleCount % logAdjustmentsInterval) == 0 {
		for k, v := range theState.Adjustments {
			log.Printf("Product: %v number of adjustments: %v", k, len(v))
			for i, el := range v {
				log.Printf("  Adjustment: %v value: %v", i, util.JsonToString(el))
			}
		}
		// clear map
		theState.Adjustments = make(map[string][]Adjustment)
	}
}
