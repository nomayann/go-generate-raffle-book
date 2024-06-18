package parameters

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadParameters(path string) (Parameters, error) {
	paramContent, err := os.ReadFile(path)
	if err != nil {
		return Parameters{}, fmt.Errorf("erreur d'ouverture du fichier de paramètres (%s)", err.Error())
	}
	var parameters Parameters
	err = json.Unmarshal(paramContent, &parameters)
	if err != nil {
		return Parameters{}, fmt.Errorf("cannot parse json parameters (%s)", err.Error())
	}

	return parameters, nil
}

func (p *Parameters) Customize(c Customs) {
	if len(c.LeftTitle) > 0 {
		p.TemplateDirectives.LeftTitle.Texts = c.LeftTitle
	}
	if len(c.RightTile) > 0 {
		p.TemplateDirectives.RightTitle.Texts = c.RightTile
	}
	if len(c.EventDescription) > 0 {
		p.TemplateDirectives.EventDescription.Texts = c.EventDescription
	}
	if len(c.PrizeTitle) > 0 {
		p.TemplateDirectives.PrizeTitle.Texts = c.PrizeTitle
	}
	if len(c.PrizeList) > 0 {
		p.TemplateDirectives.PrizeList.Texts = c.PrizeList
	}
	if c.Price != "" {
		p.TemplateDirectives.Price.Text = c.Price
	}
	if c.Buyer != "" {
		p.TemplateDirectives.Buyer.Text = c.Buyer
	}
	if c.Seller != "" {
		p.TemplateDirectives.Seller.Text = c.Seller
	}
}

type Parameters struct {
	GeneralDirectives  GeneralDirectives  `json:"generalDirectives"`
	TemplateDirectives TemplateDirectives `json:"templateDirectives"`
}

func LoadParametersWithCustom(paramPath string, customPath string) (Parameters, error) {
	parameters, err := LoadParameters("assets/parameters.json")
	if err != nil {
		return parameters, err
	}
	customs, err := LoadCustoms(customPath)
	if err != nil {
		return parameters, err
	}

	parameters.Customize(customs)

	return parameters, nil
}