package main

import "github.com/uptrace/bun"

type TyreModel struct {
	bun.BaseModel `bun:"table:tyres,alias:u"`

	Sku            string   `bun:"sku,pk" json:"sku"`
	Brand_sku      string   `bun:"brand_sku" json:"brandSku"`
	Brand          string   `bun:"brand" json:"brand"`
	Model          string   `bun:"model" json:"model"`
	Title          string   `bun:"title" json:"title"`
	Season         string   `bun:"season" json:"season"`
	Diameter       string   `bun:"diameter" json:"diameter"`
	Width          float32  `bun:"width" json:"width"`
	Profile        float32  `bun:"profile" json:"profile"`
	LoadIndex      string   `bun:"loadIndex" json:"loadIndex"`
	SpeedIndex     string   `bun:"speedIndex" json:"speedIndex"`
	Pins           bool     `bun:"pins" json:"pins"`
	RunFlat        bool     `bun:"runflat" json:"runflat"`
	Homologation   string   `bun:"homologation" json:"homologation"`
	ExtraLoad      bool     `bun:"extraLoad" json:"extraLoad"`
	Acoustic       bool     `bun:"acoustic" json:"acoustic"`
	Seal           bool     `bun:"seal" json:"seal"`
	Sale           bool     `bun:"sale" json:"sale"`
	ProductionYear string   `bun:"productionYear" json:"productionYear"`
	Price          int      `bun:"price" json:"price"`
	PriceRetail    int      `bun:"priceRetail" json:"priceRetail"`
	PriceMsrp      int      `bun:"priceMsrp" json:"priceMsrp"`
	PhotoUrl       string   `bun:"photoUrl" json:"photoUrl"`
	AmountTotal    int      `bun:"amountTotal" json:"amountTotal"`
	AmountLocal    int      `bun:"amountLocal" json:"amountLocal"`
	AmountDetailed []string `bun:"amountDetailed,array" json:"amountDetailed"`
	MarginPrice    int      `bun:"marginPrice" json:"marginPrice"`
}
type TyreRequestModel struct {
	Meta  Meta        `json:"meta"`
	Shops []string    `json:"shops"`
	Tyres []TyreModel `json:"tyre"`
}

type Meta struct {
	Created   string `json:"created"`
	Generator string `json:"generator"`
}

func main() {

}
