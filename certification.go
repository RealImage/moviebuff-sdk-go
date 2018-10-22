package moviebuff

//Certification has a country-specific movie certification
type Certification struct {
	//Indicates whether this certification represents safe for children
	ChildSafe bool `json:"childSafe"`
	// UUID of this certification
	UUID string `json:"uuid"`
	// Readable code for this certification
	Code string `json:"code"`
	// Country of this certification
	Country Country `json:"country"`
}

//Country has country information, as available on Qube Wire Cinemas
type Country struct {
	//Country name as available on Qube Wire Cinemas
	Name string `json:"name"`
	//ISO 2-digit Country Code
	Code string `json:"code"`
	//Country UUID as available on Qube Wire Cinemas
	UUID string `json:"uuid"`
}
