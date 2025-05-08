db.runCommand({
"find": "values",
"filter": {
"v": {
"$ne": Double(42)
}
},
"sort": {
"_id": 1
}
});
