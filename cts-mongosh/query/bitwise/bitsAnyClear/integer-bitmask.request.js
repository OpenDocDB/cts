db.runCommand({
"find": "values",
"filter": {
"v": {
"$bitsAnyClear": 6
}
},
"sort": {
"_id": 1
}
});
