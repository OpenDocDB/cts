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
"date"
]
}
}
},
{
"$project": {
"_id": 1,
"sum": {
"$add": [
"$v",
10
]
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
