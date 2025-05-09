response = {
"cursor": {
"firstBatch": [
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
}
],
"id": Long(0),
"ns": "db.composites"
},
"ok": Double(1)
}
