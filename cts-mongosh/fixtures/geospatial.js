db.geospatial.insertMany([
{
"_id": "point-a",
"location": {
"type": "Point",
"coordinates": [
10,
10
]
}
},
{
"_id": "point-b",
"location": {
"type": "Point",
"coordinates": [
20,
20
]
}
},
{
"_id": "point-c",
"location": {
"type": "Point",
"coordinates": [
30,
30
]
}
},
{
"_id": "zero",
"location": {
"type": "Point",
"coordinates": [
10,
10
]
}
},
{
"_id": "polygon",
"location": {
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
]);
