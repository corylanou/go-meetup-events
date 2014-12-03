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

	Events []Event
	Event  struct {
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
		Events Events `json:"results"`
		Meta   map[string]interface{}
	}
)

func (e Events) Len() int           { return len(e) }
func (e Events) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Events) Less(i, j int) bool { return e[i].Yes_rsvp_count < e[j].Yes_rsvp_count }
