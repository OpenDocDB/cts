db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$lookup": {
"from": "composites",
"localField": "v",
"foreignField": "v.foo",
"as": "composite_docs",
"pipeline": [
{
"$sort": {
"_id": 1
}
}
]
}
},
{
"$sort": {
"_id": 1
}
}
],
"cursor": {

}
});
