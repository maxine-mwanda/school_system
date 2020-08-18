package main

type Student struct{
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	EyeColor string `json:"eye_color"`
	Disability string `json:"disability"`
}

type Stream struct{
	North string `json:"north"`
	South string `json:"south"`
	East string `json:"east"`
	West string `json:"west"`
}

type House struct{
	Longonot string `json:"longonot"`
	Menengai string `json:"menengai"`
	Everest string `json:"everest"`
	Kilimanjaro string `json:"kilimanjaro"`
}

type NewStudent struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
	EyeColor string `json:"eye_color"`
	Disability string `json:"disability"`
	House string `json:"house"`
	Stream string `json:"stream"`
}
type JsonStructure struct {
	Stream string `json:"stream"`
	House string `json:"house"`
	Student []Student `json:"student"`
	}




