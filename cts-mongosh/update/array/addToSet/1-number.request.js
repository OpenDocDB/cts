db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"$addToSet": {
"v": 100
}
}
}
]
});
