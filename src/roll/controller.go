package roll

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/campaign", newCampaign).Methods("POST")
	router.HandleFunc("/campaign/{name}/roll", roll).Methods("GET")

	router.HandleFunc("/selection/{name}/{option}", selection).Methods("POST")
}

func newCampaign(resp http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var jsonData CampaignDTO
	decoder.Decode(&jsonData)

	var c = Campaign{
		Name: strings.ToLower(jsonData.Name),
	}

	c.Options = make([]Selection, len(jsonData.Options))

	for i, val := range jsonData.Options {
		c.Options[i] = Selection{
			Option:   val,
			Selected: 1,
			Offered:  1,
		}
	}

	write(c)
}

func roll(resp http.ResponseWriter, req *http.Request) {

	campaign := FindCampaign(strings.ToLower(mux.Vars(req)["name"]))

	if randRoll := rand.Float64(); randRoll <= .1 {

		randOption := rand.Intn(len(campaign.Options))

		campaign.Options[randOption].Offered++

		fmt.Println("rand", randRoll)
		updateOffered(campaign.Name, campaign.Options[randOption].Option)

		val, _ := json.Marshal(campaign.Options[randOption])

		resp.Write(val)
	} else {

		selected := 0

		for i, opt := range campaign.Options {

			if opt.Selected/opt.Offered > campaign.Options[selected].Selected/campaign.Options[selected].Offered {

				selected = i
			}
		}

		campaign.Options[selected].Offered++

		fmt.Println("pull lever", randRoll)
		updateOffered(campaign.Name, campaign.Options[selected].Option)

		val, _ := json.Marshal(campaign.Options[selected])

		resp.Write(val)
	}
}

func selection(resp http.ResponseWriter, req *http.Request) {

	name := strings.ToLower(mux.Vars(req)["name"])
	option := mux.Vars(req)["option"]

	updateSelected(name, option)
}
