db.runCommand({
"find": "values",
"filter": {
"$or": [
{
"v": {
"$eq": Double(42)
}
},
{
"v": {
"$eq": Double(0)
}
}
]
},
"sort": {
"_id": 1
}
});
