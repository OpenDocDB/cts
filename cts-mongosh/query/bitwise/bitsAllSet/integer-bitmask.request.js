db.runCommand({
"find": "values",
"filter": {
"v": {
"$bitsAllSet": 6
}
},
"sort": {
"_id": 1
}
});
