response = {
"cursor": {
"firstBatch": [
{
"_id": "array",
"v": [
42
]
},
{
"_id": "array-composite",
"v": [
42.13,
"foo",
BinData(128, "KgAN"),
ObjectId("000102030405060708091011"),
true,
ISODate("2021-11-01T10:18:42.123Z"),
null,
RegExp("foo", "i"),
42,
Timestamp({t: 42, i: 13}),
Long(41),
Decimal128("3476215962376601600.461")
]
},
{
"_id": "array-composite-reverse",
"v": [
Decimal128("3476215962376601600.461"),
Long(41),
Timestamp({t: 42, i: 13}),
42,
RegExp("foo", "i"),
null,
ISODate("2021-11-01T10:18:42.123Z"),
true,
ObjectId("000102030405060708091011"),
BinData(128, "KgAN"),
"foo",
42.13
]
},
{
"_id": "array-double-big",
"_comment": "The array contains the value x such that x+1 == x.",
"v": [
Double(2305843009213694000)
]
},
{
"_id": "array-double-max-integer",
"_comment": "The array contains the largest integer value representable as a double.",
"v": [
Double(9007199254740991)
]
},
{
"_id": "array-numbers-asc",
"v": [
42,
Long(43),
45.5,
Decimal128("3476215962376601600.461")
]
}
],
"id": Long(0),
"ns": "db.composites"
},
"ok": Double(1)
}
