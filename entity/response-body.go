package entity

type ResponseBody struct {
	Info    Info        `json:"info"`
	Results []Character `json:"results"`
}
