db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$group": {
"_id": {
"$type": "$v"
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
