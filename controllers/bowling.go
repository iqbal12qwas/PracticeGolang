package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"practice/entity"
)

//Model / Struct BowlingContainer
type BowlingContainer struct {
	Nama   string `json:"nama"`
	Isi    int    `json:"isi"`
	Status string `json:"status"`
}

func FullFill(ball int) int {
	// Max Filling 5 Balls
	for i := 0; i < 4; i++ {
		ball++
	}
	return ball
}

// Declaration Container
var wadah1 = 0
var wadah2 = 0
var wadah3 = 0
var status1 = "Not Ready"
var status2 = "Not Ready"
var status3 = "Not Ready"

type BowlingContainers []BowlingContainer

// First Declararation
var lists = BowlingContainers{
	BowlingContainer{Nama: "Wadah 1", Isi: wadah1, Status: status1},
	BowlingContainer{Nama: "Wadah 2", Isi: wadah2, Status: status2},
	BowlingContainer{Nama: "Wadah 3", Isi: wadah3, Status: status3},
}

// For Random Choosing Container
var in = []int{1, 2, 3}

func ChooseContainer(w http.ResponseWriter, r *http.Request) {

	// Choose Random Container
	randomIndex := rand.Intn(len(in))
	pick := in[randomIndex]

	// START Container Filling

	// Filling Container 1
	if pick == 1 {
		wadah1 = FullFill(1)
		status1 = "Ready"
		var t_wadah1 *int = &wadah1
		var t_status1 *string = &status1
		for _, bowling := range lists {
			if bowling.Nama == "Wadah 1" {
				lists[0] = BowlingContainer{Nama: bowling.Nama, Isi: *t_wadah1, Status: *t_status1}
			}
		}

		// Filling Container 2
	} else if pick == 2 {
		wadah2 = FullFill(1)
		status2 = "Ready"
		var t_wadah2 *int = &wadah2
		var t_status2 *string = &status2
		for _, bowling := range lists {
			if bowling.Nama == "Wadah 2" {
				lists[1] = BowlingContainer{Nama: bowling.Nama, Isi: *t_wadah2, Status: *t_status2}
			}
		}

		// Filling Container 3
	} else {
		wadah3 = FullFill(1)
		status3 = "Ready"
		var t_wadah3 *int = &wadah3
		var t_status3 *string = &status3
		for _, bowling := range lists {
			if bowling.Nama == "Wadah 3" {
				lists[2] = BowlingContainer{Nama: bowling.Nama, Isi: *t_wadah3, Status: *t_status3}
			}
		}
	}

	// END Container Filling

	// Show Response & Parse To Json
	w.Header().Set("Content-Type", "application/json")
	resp := entity.Response{Http: "GET/200", Message: "Success", Data: lists}
	json.NewEncoder(w).Encode(resp)
}
