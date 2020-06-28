package pkg

// Fortune stores the cookies in a string array
type Fortune struct {
	Cookies []*string
	SetName *string
}

// Answer is the actual reply the Fortune server
// dishes out to get requests
type Answer struct {
	Time     string `json:"time"`
	Fortune  string `json:"fortune"`
	Hostname string `json:"hostname"`
	Set      string `json:"set"`
}
