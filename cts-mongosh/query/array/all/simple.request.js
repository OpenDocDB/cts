db.runCommand({
"find": "composites",
"filter": {
"v": {
"$all": [
42,
"foo"
]
}
},
"sort": {
"_id": 1
}
});
