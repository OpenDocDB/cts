db.runCommand({
"find": "composites",
"filter": {

},
"projection": {
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
