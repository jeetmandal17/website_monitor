package types

// This stucture is defined to save details of the website
type Website struct{
	URL string
}

// Create new instance of the website
func NewWebsite(URL string) (*Website){
	return &Website{
		URL: URL,
	}
}

// This is used to forward the response to the websites
type WebsiteResponse struct {
	URL string
	Active bool
}

// This is used to define a WebsiteResponse object
func NewResponseWebsite(URL string, Active bool) (*WebsiteResponse) {
	return &WebsiteResponse{
		URL: URL,
		Active: Active,
	}
}


