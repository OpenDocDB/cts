db.runCommand({
"find": "composites",
"filter": {
"v": {
"$elemMatch": {
"$gte": 42
}
}
},
"projection": {
"v.$": 1
},
"sort": {
"_id": 1
}
});
