package main

type (
	venue struct {
		Id        int
		zip       string
		Lon       float64
		Repinned  bool
		Name      string
		State     string
		Address_1 string
		Lat       float64
		City      string
		Country   string
	}

	group struct {
		Id        int
		Created   int
		Group_lat float64
		Name      string
		Group_lon float64
		Join_mode string
		Urlname   string
		Who       string
	}

	event struct {
		Status           string
		Visibility       string
		Maybe_rsvp_count int
		Venue            venue
		Id               string
		Utc_offset       int
		Distance         float64
		Duration         int
		Time             int64
		Waitlist_count   int
		Updated          int
		Yes_rsvp_count   int
		Created          int
		Event_url        string
		Description      string
		Name             string
		Headcount        int
		Group            group
	}

	rawData struct {
		Events []event `json:"results"`
		Meta   map[string]interface{}
	}
)
