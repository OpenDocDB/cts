db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$geoWithin": {
"$polygon": [
[
5,
5
],
[
25,
5
],
[
25,
25
],
[
5,
25
],
[
5,
5
]
]
}
}
},
"sort": {
"_id": 1
}
});
