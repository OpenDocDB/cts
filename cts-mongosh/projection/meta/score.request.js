db.runCommand({
"find": "values",
"filter": {

},
"projection": {
"score": {
"$meta": "indexKey"
}
},
"sort": {
"_id": 1
}
});
