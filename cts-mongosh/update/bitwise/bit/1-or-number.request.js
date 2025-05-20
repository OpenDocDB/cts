db.runCommand({
"update": "values",
"updates": [
{
"q": {

},
"u": {
"$bit": {
"v": {
"or": 4
}
}
}
}
]
});
