package main

type Synapse struct {
	Name  string  `json:"name"`
	Ends  [2]int  `json:"ends"`
	Style *Style  `json:"style"`
	Point *[2]int `json:"point"`
}
