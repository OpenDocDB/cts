db.runCommand({
"aggregate": "admin",
"pipeline": [
{
"$currentOp": {
"allUsers": true
}
}
],
"cursor": {

}
});
