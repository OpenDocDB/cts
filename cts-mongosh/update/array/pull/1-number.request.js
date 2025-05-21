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
"$pull": {
"v": 42
}
},
"multi": true
}
]
});
