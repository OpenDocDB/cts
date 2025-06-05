db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$addFields": {
"foo": "bar"
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
