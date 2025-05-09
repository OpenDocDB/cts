response = {
"cursor": {
"firstBatch": [
{
"_id": "double",
"v": 42.13
},
{
"_id": "double-big",
"_comment": "Value x such as x+1 == x.",
"v": Double(2305843009213694000)
},
{
"_id": "double-infinity",
"v": Double(+Infinity)
},
{
"_id": "double-max",
"_comment": "The largest positive finite value representable as a double.",
"v": Double(179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)
},
{
"_id": "double-max-integer",
"_comment": "The largest integer value representable as a double.",
"v": Double(9007199254740991)
},
{
"_id": "double-nan",
"v": Double(NaN)
},
{
"_id": "double-negative-infinity",
"v": Double(-Infinity)
},
{
"_id": "double-negative-zero",
"v": Double(-0)
},
{
"_id": "double-smallest-non-zero",
"_comment": "The smallest positive non-zero value representable as a double.",
"v": 0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005
},
{
"_id": "double-whole",
"v": Double(42)
},
{
"_id": "double-zero",
"v": Double(0)
},
{
"_id": "unset"
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
