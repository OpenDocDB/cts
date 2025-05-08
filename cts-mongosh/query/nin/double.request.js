db.runCommand({
"find": "values",
"filter": {
"v": {
"$nin": [
Double(42),
Double(0)
]
}
},
"sort": {
"_id": 1
}
});
