db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$out": "output_collection"
}
]
});
