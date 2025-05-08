db.runCommand({
"find": "composites",
"filter": {
"v.foo": {
"$exists": true
}
},
"sort": {
"_id": 1
}
});
