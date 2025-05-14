db.runCommand({
"find": "values",
"filter": {
"v": {
"$mod": [
2,
0
]
}
},
"sort": {
"_id": 1
}
});
