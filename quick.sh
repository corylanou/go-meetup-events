if [ -f ./dev.env ]; then
	source ./dev.env
fi
./curl_json.sh | ./go-meetup-events --brief=true | column -t -s,
