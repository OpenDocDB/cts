db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$geoWithin": {
"$geometry": {
"type": "Polygon",
"coordinates": [
[
[
0,
0
],
[
0,
40
],
[
40,
40
],
[
40,
0
],
[
0,
0
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
