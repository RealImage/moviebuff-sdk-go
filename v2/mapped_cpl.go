package moviebuff

// MappedCPL represents a mapped CPL (Composition Playlist) response from the API
type MappedCPL struct {
	CPLID      string `json:"cpl_id"`
	PartNumber int    `json:"part_number"`
}
