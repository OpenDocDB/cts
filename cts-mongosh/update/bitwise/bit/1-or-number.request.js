db.runCommand({
"update": "values",
"updates": [
{
"q": {
"v": {
"$type": [
"int",
"long"
]
}
},
"multi": true,
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
