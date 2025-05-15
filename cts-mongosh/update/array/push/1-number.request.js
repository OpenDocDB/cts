db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"$push": {
"v": 100
}
}
}
]
});
