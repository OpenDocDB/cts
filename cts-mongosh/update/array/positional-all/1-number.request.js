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
"$set": {
"v.$[]": 100
}
},
"multi": true
}
]
});
