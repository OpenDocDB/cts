db.runCommand({
"find": "values",
"sort": {
"_id": 1
},
"projection": {
"_id": 1,
"v": 1
}
});
