db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$unset": "v"
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
