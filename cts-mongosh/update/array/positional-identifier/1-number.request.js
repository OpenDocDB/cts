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
"v.$[element]": 100
}
},
"arrayFilters": [
{
"element": 42
}
],
"multi": true
}
]
});
