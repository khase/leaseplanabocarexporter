#!/bin/bash

value=`cat ./changed.txt`
echo $value

curl --data-urlencode "chat_id=$CHAT_ID" --data-urlencode "parse_mode=Markdown" --data-urlencode "text=$value" "https://api.telegram.org/bot$API_KEY/sendMessage"