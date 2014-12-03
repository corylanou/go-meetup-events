# Usage

## Export varialbes needed to run:

These export variables are required to run `events.sh`.

```bash
export meetup_sig=<api key>
export meetup_id=<your meetup account id>
```

## Running the script

The `events.sh` command will run a curl command to the api meetup api, and pipe it's output
to the go program that converts the json to a minimal csv file.

You can modify the url to change the categories or query options

[Meetup Documentation for Events](http://www.meetup.com/meetup_api/docs/2/open_events/)


```bash
./events.sh
```

To view this in the terminal, you can use the `column` command:

```bash
./events.sh | column -s, -t
```

And then of course you can pipe that to `more`:

```bash
./events.sh | column -s, -t | more
```


