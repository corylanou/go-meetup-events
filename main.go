package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

func main() {
	var raw rawData

	stdin := bufio.NewReader(os.Stdin)
	decoder := json.NewDecoder(stdin)
	err := decoder.Decode(&raw)
	if err != nil {
		panic(err)
	}

	fmt.Println("GroupId,GroupName,YesRSVPCount,Venue,EventUrl,City,State,Time")
	events := raw.Events
	sort.Sort(sort.Reverse(raw.Events))
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
