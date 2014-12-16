# Usage

## Export varialbes needed to run:

These export variables are required to run `events.sh`.

```bash
export meetup_sig=<api key>
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

## Categories

If you want to change the url for categories, I grabbed the latest ones.  Here is the endpoint if you want it:


[Meetup Documentation for Categories](http://www.meetup.com/meetup_api/docs/2/categories/)

```json
HTTP/1.1 200 success
{
"results": [
{
"id": 1,
"name": "Arts & Culture",
"shortname": "Arts"
},
{
"id": 2,
"name": "Career & Business",
"shortname": "Business"
},
{
"id": 3,
"name": "Cars & Motorcycles",
"shortname": "Auto"
},
{
"id": 4,
"name": "Community & Environment",
"shortname": "Community"
},
{
"id": 5,
"name": "Dancing",
"shortname": "Dancing"
},
{
"id": 6,
"name": "Education & Learning",
"shortname": "Education"
},
{
"id": 8,
"name": "Fashion & Beauty",
"shortname": "Fashion"
},
{
"id": 9,
"name": "Fitness",
"shortname": "Fitness"
},
{
"id": 10,
"name": "Food & Drink",
"shortname": "Food & Drink"
},
{
"id": 11,
"name": "Games",
"shortname": "Games"
},
{
"id": 13,
"name": "Movements & Politics",
"shortname": "Movements"
},
{
"id": 14,
"name": "Health & Wellbeing",
"shortname": "Well-being"
},
{
"id": 15,
"name": "Hobbies & Crafts",
"shortname": "Crafts"
},
{
"id": 16,
"name": "Language & Ethnic Identity",
"shortname": "Languages"
},
{
"id": 12,
"name": "LGBT",
"shortname": "LGBT"
},
{
"id": 17,
"name": "Lifestyle",
"shortname": "Lifestyle"
},
{
"id": 18,
"name": "Literature & Writing",
"shortname": "Literature"
},
{
"id": 20,
"name": "Movies & Film",
"shortname": "Films"
},
{
"id": 21,
"name": "Music",
"shortname": "Music"
},
{
"id": 22,
"name": "New Age & Spirituality",
"shortname": "Spirituality"
},
{
"id": 23,
"name": "Outdoors & Adventure",
"shortname": "Outdoors"
},
{
"id": 24,
"name": "Paranormal",
"shortname": "Paranormal"
},
{
"id": 25,
"name": "Parents & Family",
"shortname": "Moms & Dads"
},
{
"id": 26,
"name": "Pets & Animals",
"shortname": "Pets"
},
{
"id": 27,
"name": "Photography",
"shortname": "Photography"
},
{
"id": 28,
"name": "Religion & Beliefs",
"shortname": "Beliefs"
},
{
"id": 29,
"name": "Sci-Fi & Fantasy",
"shortname": "Sci fi"
},
{
"id": 30,
"name": "Singles",
"shortname": "Singles"
},
{
"id": 31,
"name": "Socializing",
"shortname": "Social"
},
{
"id": 32,
"name": "Sports & Recreation",
"shortname": "Sports"
},
{
"id": 33,
"name": "Support",
"shortname": "Support"
},
{
"id": 34,
"name": "Tech",
"shortname": "Tech"
},
{
"id": 35,
"name": "Women",
"shortname": "Women"
}
],
"meta": {
"lon": "",
"count": 33,
"signed_url": "http://secure.meetup.com/2/categories?order=shortname&desc=false&offset=0&photo-host=public&format=json&page=100&sig_id=683412&sig=b7d14a0e57bf1d77fefdeed5801428bc0ca873ec",
"link": "https://secure.meetup.com/2/categories",
"next": "",
"total_count": 33,
"url": "https://secure.meetup.com/2/categories?order=shortname&desc=false&offset=0&photo-host=public&format=json&page=100&sign=True",
"id": "",
"title": "Categories",
"updated": 1401302948000,
"description": "Returns a list of Meetup group categories",
"method": "Categories",
"lat": ""
}
}
```
