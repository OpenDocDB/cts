db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "int32"
},
"u": {
"$max": {
"v": 50
}
}
}
]
});
