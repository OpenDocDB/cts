db.runCommand({
"find": "values",
"filter": {
"$and": [
{
"v": {
"$gte": Double(0)
}
},
{
"v": {
"$lte": Double(42)
}
}
]
},
"sort": {
"_id": 1
}
});
