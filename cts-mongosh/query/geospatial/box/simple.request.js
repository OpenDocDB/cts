db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$geoWithin": {
"$box": [
[
5,
5
],
[
25,
25
]
]
}
}
},
"sort": {
"_id": 1
}
});
