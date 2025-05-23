db.runCommand({
"createIndexes": "geospatial",
"indexes": [
{
"key": {
"location": "2dsphere"
},
"name": "location_2dsphere"
}
]
});
