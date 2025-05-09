db.runCommand({
"find": "values",
"filter": {
"v": {
"$lte": Double(42)
}
},
"sort": {
"_id": 1
}
});
