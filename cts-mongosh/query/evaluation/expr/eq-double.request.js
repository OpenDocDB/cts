db.runCommand({
"find": "values",
"filter": {
"$expr": {
"$eq": [
"$v",
Double(42)
]
}
},
"sort": {
"_id": 1
}
});
