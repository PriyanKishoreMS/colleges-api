package entities

type College struct {
	Name          string `json: "name"`
	State         string `json: "state"`
	City          string `json: "city"`
	Address_line1 string `json: "address_line1"`
	Address_line2 string `json: "address_line2"`
}
