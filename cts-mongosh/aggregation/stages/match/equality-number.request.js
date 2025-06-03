db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$match": {
"v": 42
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
