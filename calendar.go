package moviebuff

//Calendar - contains the info about the country's public holidays
type Calendar struct {
	ID        string    `json:"calendarId"`
	Name      string    `json:"name"`
	Holidays  []Holiday `json:"holidays"`
	SyncToken string    `json:"syncToken"`
	TimeZone  string    `json:"timeZone"`
}

//Holiday - contains id, date and name of the holiday
type Holiday struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}
