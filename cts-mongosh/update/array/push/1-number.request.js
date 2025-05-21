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
"$push": {
"v": 100
}
},
"multi": true
}
]
});
