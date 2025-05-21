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
"$pullAll": {
"v": [
42,
100
]
}
},
"multi": true
}
]
});
