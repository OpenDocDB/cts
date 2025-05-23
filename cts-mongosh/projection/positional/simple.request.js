db.runCommand({
"find": "composites",
"filter": {
"v": {
"$elemMatch": {
"$exists": true
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
