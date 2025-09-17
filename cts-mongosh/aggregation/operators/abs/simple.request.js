db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$match": {
"v": {
"$type": [
"int",
"long",
"double",
"decimal",
"null"
],
"$gt": Long(-9223372036854775808)
}
}
},
{
"$project": {
"_id": 1,
"abs": {
"$abs": "$v"
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
