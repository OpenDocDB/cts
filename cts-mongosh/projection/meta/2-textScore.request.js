db.runCommand({
"find": "values",
"filter": {
"$text": {
"$search": "foo"
}
},
"projection": {
"score": {
"$meta": "textScore"
}
},
"sort": {
"_id": 1
}
});
