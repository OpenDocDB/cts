db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$bucket": {
"groupBy": "$v",
"boundaries": [
0,
41,
100
],
"default": "Other",
"output": {
"count": {
"$sum": 1
}
}
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
