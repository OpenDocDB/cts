db.runCommand({
"find": "values",
"filter": {
"v": {
"$type": "double"
}
},
"sort": {
"_id": 1
}
});
