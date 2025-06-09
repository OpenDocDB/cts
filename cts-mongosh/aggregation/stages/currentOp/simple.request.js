db.runCommand({
"aggregate": 1,
"pipeline": [
{
"$currentOp": {

}
},
{
"$project": {
"type": 1,
"op": 1,
"ns": 1,
"command": 1
}
}
],
"cursor": {

}
});
