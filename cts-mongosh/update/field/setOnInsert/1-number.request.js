db.runCommand({
"update": "values",
"updates": [
{
"q": {
"_id": "new-document"
},
"u": {
"$setOnInsert": {
"v": 42
}
},
"upsert": true
}
]
});
