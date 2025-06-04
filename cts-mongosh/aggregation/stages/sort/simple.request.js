db.runCommand({
"aggregate": "values",
"pipeline": [
{
"$sort": {
"_id": 1
}
}
],
"cursor": {

}
});
