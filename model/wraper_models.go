package model

type Link struct {
	Url        string `json:"url"`
	Accessible bool   `json:"accessible"`
}

type HtmlDocument struct {
	HTMLVersion   string `json:"html_version"`
	Title         string `json:"title"`
	LoginExist    bool   `json:"login_exist"`
	H1Count       int    `json:"h1"`
	H2Count       int    `json:"h2"`
	H3Count       int    `json:"h3"`
	H4Count       int    `json:"h4"`
	H5Count       int    `json:"h5"`
	H6Count       int    `json:"h6"`
	InternalLinks []Link `json:"internal_links"`
	ExternalLinks []Link `json:"external_links"`
	Host          string `json:"-"`
}

type Request struct {
	URL string `json:"url"`
}
