db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "int32"
},
"u": {
"$min": {
"v": 20
}
}
}
]
});
