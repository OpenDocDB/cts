db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "int32"
},
"u": {
"$mul": {
"v": 2
}
}
}
]
});
