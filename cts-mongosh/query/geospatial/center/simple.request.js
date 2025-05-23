db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$geoWithin": {
"$center": [
[
15,
15
],
10
]
}
}
},
"sort": {
"_id": 1
}
});
