# Usage

## Export varialbes needed to run:

`bash
export meetup_sig=<api key>
export meetup_id=<your meetup account id>
`

## Running the curl command

The `events.sh` command will run a curl command to the api.  It requires the above variables to be exported to run.

You can modify the url to change the categories or query options

[Meetup Documentation for Events](http://www.meetup.com/meetup_api/docs/2/open_events/)


`bash
./events.sh
`

## Converting to CSV

`bash
./events.sh | ./go-meetup-events
`
