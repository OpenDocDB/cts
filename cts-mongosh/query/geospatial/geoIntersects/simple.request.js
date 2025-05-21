db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$geoIntersects": {
"$geometry": {
"type": "Polygon",
"coordinates": [
[
[
5,
5
],
[
5,
15
],
[
15,
15
],
[
15,
5
],
[
5,
5
]
]
]
}
}
}
},
"sort": {
"_id": 1
}
});
