db.runCommand({
"find": "values",
"filter": {
"v": {
"$bitsAnySet": 6
}
},
"sort": {
"_id": 1
}
});
