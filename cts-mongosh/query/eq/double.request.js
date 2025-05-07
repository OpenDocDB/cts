db.runCommand({
"find": "values",
"filter": {
"v": {
"$eq": Double(0)
}
},
"sort": {
"_id": 1
}
});
