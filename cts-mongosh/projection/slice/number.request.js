db.runCommand({
"find": "composites",
"filter": {

},
"projection": {
"v": {
"$slice": 2
}
},
"sort": {
"_id": 1
}
});
