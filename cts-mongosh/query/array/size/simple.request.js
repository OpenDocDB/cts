db.runCommand({
"find": "composites",
"filter": {
"v": {
"$size": 3
}
},
"sort": {
"_id": 1
}
});
