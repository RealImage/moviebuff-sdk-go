package moviebuff

type MappedCPL struct {
	ID               int            `json:"id"`
	UUID             string         `json:"uuid"`
	ContentTitleText string         `json:"content_title_text"`
	Movie            MappedCPLMovie `json:"movie"`
}

type MappedCPLMovie struct {
	ID   int            `json:"id"`
	Name string         `json:"name"`
	UUID string         `json:"uuid"`
	Part *MappedCPLPart `json:"part,omitempty"`
}

type MappedCPLPart struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
	Name string `json:"name"`
}
