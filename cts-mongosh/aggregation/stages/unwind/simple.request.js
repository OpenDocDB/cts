db.runCommand({
"aggregate": "composites",
"pipeline": [
{
"$unwind": {
"path": "$v"
}
},
{
"$sort": {
"_id": 1,
"v": 1
}
}
],
"cursor": {

}
});
