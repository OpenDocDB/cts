db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "int32"
},
"u": {
"$inc": {
"v": 10
}
}
}
]
});
