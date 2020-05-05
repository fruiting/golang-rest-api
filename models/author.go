package models

type Author struct {
	Firstname string `id:"first_name" json:"first_name"`
	Lastname  string `id:"last_name" json:"last_name"`
}
