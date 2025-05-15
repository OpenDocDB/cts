db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"v.$[element]": 100
},
"arrayFilters": [
{
"element": 42
}
]
}
]
});
