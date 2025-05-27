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
"$addToSet": {
"v": 100
}
},
"multi": true
}
]
});
