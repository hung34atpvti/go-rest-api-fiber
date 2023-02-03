package organization

//Organization struct
type Organization struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Organizations struct {
	Organizations []Organization `json:"organizations"`
}
