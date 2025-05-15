db.runCommand({
"update": "composites",
"updates": [
{
"q": {
"_id": "array"
},
"u": {
"$pullAll": {
"v": [
42,
100
]
}
}
}
]
});
