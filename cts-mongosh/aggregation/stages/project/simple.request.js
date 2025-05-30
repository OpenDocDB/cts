db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$project": {
"_id": 1,
"v": 1
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
