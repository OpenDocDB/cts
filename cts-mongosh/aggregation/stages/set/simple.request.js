db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$set": {
"new": "foobar"
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
