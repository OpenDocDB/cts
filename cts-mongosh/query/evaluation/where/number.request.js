db.runCommand({
"find": "values",
"filter": {
"$where": "this.v == 42"
},
"sort": {
"_id": 1
}
});
