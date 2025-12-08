package moviebuff

// MappedCPL represents a mapped CPL (Composition Playlist) response from the API
type MappedCPL struct {
	// ID is the unique identifier of the mapped CPL
	ID int `json:"id"`

	// UUID is the universally unique identifier of the mapped CPL
	UUID string `json:"uuid"`

	// ContentTitleText is the content title text from the CPL
	ContentTitleText string `json:"content_title_text"`

	// Movie contains the associated movie information
	Movie MappedCPLMovie `json:"movie"`
}

// MappedCPLMovie represents the movie information in a mapped CPL response
type MappedCPLMovie struct {
	// ID is the unique identifier of the movie
	ID int `json:"id"`

	// Name is the name of the movie
	Name string `json:"name"`

	// UUID is the universally unique identifier of the movie
	UUID string `json:"uuid"`

	// Part contains the part information if the movie has multiple parts
	Part *MappedCPLPart `json:"part,omitempty"`
}

// MappedCPLPart represents the part information in a mapped CPL response
type MappedCPLPart struct {
	// ID is the unique identifier of the part
	ID int `json:"id"`

	// UUID is the universally unique identifier of the part
	UUID string `json:"uuid"`

	// Name is the name of the part (e.g., "Part 2")
	Name string `json:"name"`
}
