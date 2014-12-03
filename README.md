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

If you want to do a "brief" layout, you can issue this set of commands:

```bash
./curl_json.sh | ./go-meetup-events --brief | column -s, -t
```

That will produce this type of output:

```bash
GroupName                                                       RSVP       Venue                             City           State                          Time
"Analyze Boulder"                                               150        "Galvanize Boulder"               "Boulder"      "CO"                           "2014-12-04 18:00:00 -0700 MST"
"Data Science & Business Analytics"                             89         "ATLAS Building"                  "Boulder"      "CO"                           "2014-12-03 18:00:00 -0700 MST"
"SheSays Boulder"                                               82         ""                                ""             ""                             "2014-12-03 18:00:00 -0700 MST"
"CoFoundersLab Matchup Denver"                                  50         "Galvanize"                       "Denver"       "CO"                           "2014-12-03 18:00:00 -0700 MST"
"Women Who Code (Boulder/Denver)"                               48         "SendGrid Denver Office"          "Denver"       "CO"                           "2014-12-03 18:00:00 -0700 MST"
"MJ Startup Denver"                                             45         ""                                ""             ""                             "2014-12-04 19:00:00 -0700 MST"
"Mobile Monday Colorado"                                        38         ""                                ""             ""                             "2014-12-08 18:00:00 -0700 MST"
"Sophos User Group - Colorado"                                  28         "Ale House at Amato's"            "Denver"       "Co"                           "2014-12-03 16:00:00 -0700 MST"
"Colorado independent Game Developers Association (CiGDA)"      24         "Panera Bread"                    "Superior"     "CO"                           "2014-12-05 19:00:00 -0700 MST"
"Denver Visual Studio User Group"                               18         "Microsoft"                       "Denver"       "CO"                           "2014-12-08 17:30:00 -0700 MST"
"Women Who Code (Boulder/Denver)"                               18         "Boulder Cafe"                    "Boulder"      "CO"                           "2014-12-04 18:00:00 -0700 MST"
"Turing Community Events"                                       16         "Turing"                          "Denver"       "CO"                           "2014-12-04 17:30:00 -0700 MST"
"Maptime Boulder"                                               12         "Discovery Learning Center"       "Boulder"      "CO"                           "2014-12-03 18:30:00 -0700 MST"
"Women Who Code (Boulder/Denver)"                               12         "Galvanize"                       "Denver"       "CO"                           "2014-12-04 18:00:00 -0700 MST"
"Colorado Bitcoin Society"                                      11         "Denver BitCoin Center"           "Denver"       "CO"                           "2014-12-05 09:00:00 -0700 MST"
"Girl Develop It Boulder/Denver"                                10         "SendGrid Boulder Office"         "Boulder"      "CO"                           "2014-12-06 13:00:00 -0700 MST"
"Loveland CO Creatorspace"                                      10         "Loveland CreatorSpace"           "Loveland"     ""                             "2014-12-07 13:00:00 -0700 MST"
"ScrumMasters Guild"                                            9          "Breckenridge Brewery"            "Denver"       "CO"                           "2014-12-03 17:00:00 -0700 MST"
"Boulder Haskell Programmers"                                   9          "Simple Energy"                   "Boulder"      "CO"                           "2014-12-03 18:30:00 -0700 MST"
"Colorado Bitcoin Society"                                      9          "Denver BitCoin Center"           "Denver"       "CO"                           "2014-12-04 18:30:00 -0700 MST"
"Startup Longmont"                                              8          "300 Suns Brewing"                "Longmont"     "CO"                           "2014-12-08 17:00:00 -0700 MST"
"Colorado Bitcoin Society"                                      8          "Southern Hospitality"            "Denver"       "CO"                           "2014-12-08 18:00:00 -0700 MST"
"Big Data for Business"                                         8          "Metropoint"                      "Centennial"   "CO"                           "2014-12-04 18:30:00 -0700 MST"
"Colorado Experience Leaders (CX/UX)"                           8          "Paxia Restaurant"                "Denver"       "CO"                           "2014-12-03 18:30:00 -0700 MST"
"Denver Droids (aka Denver Android Developers Group)"           7          "Panera Bread"                    "Denver"       "CO"                           "2014-12-03 12:00:00 -0700 MST"
"DTC Ruby User Group"                                           7          "Hyatt House Denver Tech Center"  "Englewood"    "CO"                           "2014-12-03 18:15:00 -0700 MST"
"The Denver/Boulder Linux Meetup Group"                         7          "Rock Bottom Brewery"             "Westminster"  "CO"                           "2014-12-04 18:30:00 -0700 MST"
"Frontrange Monitoring Alerting& Uptime"                        7          "VictorOps Mountain Basecamp"     "Boulder"      "CO"                           "2014-12-04 18:30:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           6          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-07 14:00:00 -0700 MST"
"Boulder Music Tech Hackers"                                    6          "Applied Trust"                   "Boulder"      "CO"                           "2014-12-04 18:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           6          "TinkerMill Downtown"             "Longmont"     "CO"                           "2014-12-04 19:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           6          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-03 19:00:00 -0700 MST"
"Loveland CO Creatorspace"                                      6          "Loveland CreatorSpace"           "Loveland"     ""                             "2014-12-03 16:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           5          "TinkerMill Downtown"             "Longmont"     "CO"                           "2014-12-05 18:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           5          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-08 18:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           5          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-04 18:00:00 -0700 MST"
"Anschutz Medical Campus Digital Health Group"                  5          "COPIC"                           "Denver"       "CO"                           "2014-12-04 07:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           5          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-07 10:00:00 -0700 MST"
"Denver Bitcoin"                                                5          "Denver BitCoin Center"           "Unit # 5B"    "CO"                           "2014-12-04 18:30:00 -0700 MST"
"Denver Bitcoin"                                                5          "Denver BitCoin Center"           "Unit # 5B"    "CO"                           "2014-12-05 09:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           4          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-07 12:00:00 -0700 MST"
"STEM Enthusiasts"                                              4          ""                                ""             ""                             "2014-12-06 11:30:00 -0700 MST"
"Denver Mobile .NET Developers Group"                           4          "Tack Mobile"                     "Denver"       "CO"                           "2014-12-08 18:00:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           3          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-06 09:30:00 -0700 MST"
"TinkerMill: The Longmont Makerspace"                           3          "TinkerMill on Delaware Pl."      "Longmont"     "CO"                           "2014-12-07 13:00:00 -0700 MST"
"Denver Juniper Users Group"                                    3          ""                                ""             ""                             "2014-12-04 18:00:00 -0700 MST"
"Loveland CO Creatorspace"                                      3          "Loveland CreatorSpace"           "Loveland"     "CO"                           "2014-12-07 08:00:00 -0700 MST"
"Denver Area Access Users Group"                                3          "Microsoft Store"                 "Lone Tree"    "CO"                           "2014-12-03 18:30:00 -0700 MST"
"Denver Cryptocurrency Meetup"                                  3          ""                                ""             ""                             "2014-12-03 18:00:00 -0700 MST"
```

