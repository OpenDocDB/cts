db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"v": {
"$type": "array"
}
},
"u": {
"$pop": {
"v": 1
}
},
"multi": true
}
]
});
