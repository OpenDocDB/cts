db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array-documents"
},
"u": {
"$pop": {
"v": 1
}
}
}
]
});
