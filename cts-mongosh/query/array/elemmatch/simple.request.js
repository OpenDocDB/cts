db.runCommand({
"find": "composites",
"filter": {
"v": {
"$elemMatch": {
"$gte": 42
}
}
},
"sort": {
"_id": 1
}
});
