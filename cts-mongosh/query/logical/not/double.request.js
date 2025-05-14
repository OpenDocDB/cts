db.runCommand({
"find": "values",
"filter": {
"v": {
"$not": {
"$gte": Double(42)
}
}
},
"sort": {
"_id": 1
}
});
