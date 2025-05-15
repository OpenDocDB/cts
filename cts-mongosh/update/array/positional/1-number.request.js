db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"v.$": 100
}
}
]
});
