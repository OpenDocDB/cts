db.runCommand({
"find": "composites",
"filter": {
"_id": "array"
},
"sort": {
"_id": 1
}
});
