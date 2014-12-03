package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {

	var brief bool
	flag.BoolVar(&brief, "brief", false, "specify brief output.  defaults to false.")
	flag.Parse()

	var raw rawData

	stdin := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(stdin)
	err := decoder.Decode(&raw)
	if err != nil {
		panic(err)
	}

	events := raw.Events
	sort.Sort(sort.Reverse(raw.Events))
	if brief {
		Brief(events)
		return
	}
	Verbose(events)
}

func Verbose(events Events) {
	fmt.Println("GroupId,GroupName,RSVP,Venue,EventUrl,City,State,Time")
	for _, e := range events {
		fmt.Printf(
			"%q,%q,%d,%q,%q,%q,%q,%q\n",
			e.Id,
			e.Group.Name,
			e.Yes_rsvp_count,
			e.Venue.Name,
			e.Event_url,
			e.Venue.City,
			e.Venue.State,
			time.Unix(0, e.Time*int64(time.Millisecond)),
		)
	}
}

func Brief(events Events) {
	fmt.Println("GroupName,RSVP,Venue,City,State,Time")
	for _, e := range events {
		fmt.Printf(
			"%q,%d,%q,%q,%q,%q\n",
			e.Group.Name,
			e.Yes_rsvp_count,
			e.Venue.Name,
			e.Venue.City,
			e.Venue.State,
			time.Unix(0, e.Time*int64(time.Millisecond)),
		)
	}
}
