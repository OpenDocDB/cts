db.runCommand({
"find": "values",
"filter": {
"v": {
"$regex": "foo",
"$options": "i"
}
},
"sort": {
"_id": 1
}
});
