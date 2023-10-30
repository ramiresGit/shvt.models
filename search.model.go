package shvtmodels

type SearchModel struct {
	Brand    string  `json:"brand"`
	Season   string  `json:"season"`
	Diameter string  `json:"diameter"`
	Width    float32 `json:"width"`
	Profile  float32 `json:"profile"`
	RunFlat  bool    `json:"runflat"`
	Seal     bool    `json:"seal"`
	Acoustic bool    `json:"acoustic"`
}
