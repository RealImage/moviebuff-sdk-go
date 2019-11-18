package moviebuff

import v2 "github.com/RealImage/moviebuff-sdk-go/v2"

// Certification contains a readable code as well as UUID along with whether the certification indicates that the movie is safe for children and the country it is applicable for.
// A movie has multiple certifications, one for each country it is released in.
type Certification = v2.Certification

//Country has country information as available on Qube Wire Cinemas
type Country = v2.Country
