#!/bin/sh

command -v curl >/dev/null 2>&1 || { echo >&2 "I require curl but it's not installed.  Aborting."; exit 1; }

: ${HIPCHAT_ROOM:?"Need to set HIPCHAT_ROOM non-empty"}
: ${HIPCHAT_ADMIN_TOKEN:?"Need to set HIPCHAT_ADMIN_TOKEN non-empty"}

case "$1" in
  --test) 
    echo "Write a message to test webhook"
    while true
    do

      read -p ">" message

      curl -s --fail -H "Content-Type: application/json" -d "{\"message\":\"$message\"}" \
        https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/notification?auth_token=$HIPCHAT_ADMIN_TOKEN
      
    done
  ;;

  --create)
    : ${FORWARD_APP_URL:?"Need to set FORWARD_APP_URL non-empty"}
    curl -H "Content-Type: application/json" -d "{\"url\":\"$FORWARD_APP_URL\",\"event\": \"room_notification\", \"name\": \"$HIPCHAT_ROOM\" }" \
      https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/webhook?auth_token=$HIPCHAT_ADMIN_TOKEN
  ;;

  --check)
    curl -s https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/webhook?auth_token=$HIPCHAT_ADMIN_TOKEN
  ;;

  --clear)
    webhooks=$(curl -s --fail https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/webhook?auth_token=$HIPCHAT_ADMIN_TOKEN)
    if [ $? = 0 ]
    then
      ids=`echo $webhooks | tr ',' '\n'  | sed  '/id/!d' | sed -E 's/.*id: ([0-9]+)/\1/'`
      for id in $ids
      do
        curl -s --fail -X DELETE https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/webhook/$id?auth_token=$HIPCHAT_ADMIN_TOKEN
      done
      curl -s https://api.hipchat.com/v2/room/$HIPCHAT_ROOM/webhook?auth_token=$HIPCHAT_ADMIN_TOKEN
    fi
  ;;
esac

