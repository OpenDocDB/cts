db.runCommand({
"find": "values",
"filter": {
"_id": "int32"
},
"sort": {
"_id": 1
}
});
