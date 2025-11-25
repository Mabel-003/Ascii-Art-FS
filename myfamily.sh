curl -s https://acad.learn2earn.ng/assets/superhero/all.json \
| jq -r --arg id "$HERO_ID" '.[] | select(.id == ($id|tonumber)) | .connections.relatives' \
| awk '{printf "%s\\n", $0}' \
| sed 's/\\n$//'


