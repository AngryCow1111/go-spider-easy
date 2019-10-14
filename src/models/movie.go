package models

type Movie struct {
	ID           int32
	Director     string
	Name         string
	Year         string
	Scriptwriter string
	Staring      string
	MovieType    string
	AddDate      string
}
type Movies struct {
	Items []*Movie
}
