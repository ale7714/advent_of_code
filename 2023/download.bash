#!/bin/bash
if [ -z "$1" ]
then
    echo "Usage: download.bash <day number>"
    exit 1
fi

day=$1
secret_cookie=`cat ./secrets.txt` # Alternatively, you can directly copy and past it here. 

mkdir -p day${day}

curl -v "https://adventofcode.com/2023/day/${day}/input" \
  -H 'authority: adventofcode.com' \
  -H 'accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7' \
  -H 'accept-language: en-US,en;q=0.9' \
  -H 'cache-control: max-age=0' \
  -H "cookie: ${secret_cookie}" \
  -H "referer: https://adventofcode.com/2023/day/${day}" \
  -H 'sec-ch-ua: "Google Chrome";v="119", "Chromium";v="119", "Not?A_Brand";v="24"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "Linux"' \
  -H 'sec-fetch-dest: document' \
  -H 'sec-fetch-mode: navigate' \
  -H 'sec-fetch-site: same-origin' \
  -H 'sec-fetch-user: ?1' \
  -H 'upgrade-insecure-requests: 1' \
  -H 'user-agent: Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36' \
  --compressed > day${day}/inp.txt
  
cp template.go day${day}/main.go
