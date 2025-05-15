db.runCommand({
"find": "composites",
"filter": {
"_id": "array-documents"
},
"sort": {
"_id": 1
}
});
