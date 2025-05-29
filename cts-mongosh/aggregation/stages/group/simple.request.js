db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$group": {
"_id": {
"$type": "$v"
}
}
}
]
});
