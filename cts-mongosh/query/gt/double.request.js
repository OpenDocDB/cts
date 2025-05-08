db.runCommand({
"find": "values",
"filter": {
"v": {
"$gt": Double(42)
}
},
"sort": {
"_id": 1
}
});
