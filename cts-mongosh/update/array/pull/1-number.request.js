db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"$pull": {
"v": 42
}
}
}
]
});
