db.runCommand({
"find": "values",
"filter": {
"v": {
"$bitsAllClear": 6
}
},
"sort": {
"_id": 1
}
});
