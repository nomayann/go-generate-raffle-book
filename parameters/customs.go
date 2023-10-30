package parameters

import (
	"encoding/json"
	"fmt"
	"os"
)

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
