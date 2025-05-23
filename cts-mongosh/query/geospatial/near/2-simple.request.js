db.runCommand({
"find": "geospatial",
"filter": {
"location": {
"$near": {
"$geometry": {
"type": "Point",
"coordinates": [
15,
15
]
},
"$maxDistance": 2000000
}
}
},
"sort": {
"_id": 1
}
});
