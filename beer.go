package main

import "time"

type Beer struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Abv	  float32   `json:"abv"`
	Tapped    time.Time `json:"tapped"`
}

type Beers []Beer
