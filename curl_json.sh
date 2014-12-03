page=150
url="http://api.meetup.com/2/open_events.json/?status=upcoming&radius=50&category=34&page=$page&and_text=False&limited_events=False&desc=False&offset=0&photo-host=public&format=json&lat=39.7392&page=50&lon=-104.9847&&key=$meetup_sig"

curl $url 
