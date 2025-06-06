response = {
"cursor": {
"firstBatch": [
{
"type": "op",
"op": "command",
"ns": "admin.$cmd.aggregate",
"command": {
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

},
"$db": "admin"
}
},
{
"type": "op",
"op": "none",
"ns": "",
"command": {

}
},
{
"type": "op",
"op": "none",
"ns": "",
"command": {

}
}
],
"id": Long(0),
"ns": "admin.$cmd.aggregate"
},
"ok": Double(1)
}
