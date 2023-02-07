package entity

//Organization struct
type Organization struct {
	Id          int    `json:"id" query:"id"`
	Name        string `json:"name" query:"name"`
	Description string `json:"description" query:"description"`
}

type Organizations struct {
	Organizations []Organization `json:"organizations"`
}
