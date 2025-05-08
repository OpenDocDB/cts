db.runCommand({
"find": "values",
"filter": {
"v": {
"$gte": Double(42)
}
},
"sort": {
"_id": 1
}
});
