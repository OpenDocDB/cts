db.runCommand({
"find": "values",
"filter": {
"v": {
"$lt": Double(42)
}
},
"sort": {
"_id": 1
}
});
