db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"v": 42
},
"u": {
"$set": {
"v.$": 100
}
},
"multi": true
}
]
});
