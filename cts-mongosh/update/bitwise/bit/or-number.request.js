db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "int32"
},
"u": {
"$bit": {
"v": {
"or": 4
}
}
}
}
]
});
