package parameters

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customs struct {
	LeftTitle        []string `json:"left_title,omitempty"`
	RightTile        []string `json:"right_title,omitempty"`
	EventDescription []string `json:"event_description,omitempty"`
	PrizeTitle       []string `json:"prize_title,omitempty"`
	PrizeList        []string `json:"prize_list,omitempty"`
	Price            string   `json:"price,omitempty"`
	Buyer            string   `json:"buyer,omitempty"`
	Seller           string   `json:"seller,omitempty"`
}

func LoadCustoms(path string) (Customs, error) {
	if _, err := os.Stat(path); err != nil {
		fmt.Printf("no %s file exists. Generating without customization\n", path)

		return Customs{}, nil
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return Customs{}, fmt.Errorf("erreur d'ouverture du fichier de customizations (%s)", err.Error())
	}
	var customs Customs
	err = json.Unmarshal(content, &customs)
	if err != nil {
		return Customs{}, fmt.Errorf("fichier json customs mal formaté (%s)", err.Error())
	}

	return customs, nil
}

func CreateBoilerPlateCustoms() Customs {
	return Customs{
		LeftTitle:        []string{"left title", "next"},
		RightTile:        []string{"right title", "next"},
		EventDescription: []string{"event description", "next"},
		PrizeTitle:       []string{"prize title"},
		PrizeList:        []string{"* 1st prize", "* 2nd prize"},
		Buyer:            "Buyer name",
		Seller:           "Seller name",
		Price:            "1€",
	}
}

