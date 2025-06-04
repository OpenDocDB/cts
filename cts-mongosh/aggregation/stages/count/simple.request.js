db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$count": "total"
}
],
"cursor": {

}
});
