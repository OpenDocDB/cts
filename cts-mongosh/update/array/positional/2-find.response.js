response = {
"cursor": {
"firstBatch": [
{
"_id": "array",
"v": [
100
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
100,
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
100,
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
"_id": "array-document-empty",
"v": [
{

}
]
},
{
"_id": "array-documents",
"_comment": "An array of documents, add more documents if necessary.",
"v": [
{
"foo": 42
},
{
"foo": 44
},
{
"bar": 42
}
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
"_id": "array-double-desc",
"v": [
Double(40),
Double(15),
Double(10)
]
},
{
"_id": "array-double-duplicate",
"_comment": "The first and the second elements are the same value.",
"v": [
Double(10),
Double(10),
Double(20)
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
"_id": "array-empty",
"v": [

]
},
{
"_id": "array-nested",
"v": [
[
"foo"
]
]
},
{
"_id": "array-null",
"v": [
null
]
},
{
"_id": "array-numbers-asc",
"v": [
100,
Long(43),
45.5,
Decimal128("3476215962376601600.461")
]
},
{
"_id": "document",
"v": {
"foo": 42
}
},
{
"_id": "document-composite",
"v": {
"foo": 42,
"42": "foo",
"array": [
42,
"foo",
null
]
}
},
{
"_id": "document-composite-reverse",
"v": {
"array": [
42,
"foo",
null
],
"42": "foo",
"foo": 42
}
},
{
"_id": "document-empty",
"v": {

}
},
{
"_id": "document-nested",
"v": {
"foo": {
"document": 12
}
}
},
{
"_id": "document-null",
"v": {
"foo": null
}
}
],
"id": Long(0),
"ns": "db.composites"
},
"ok": Double(1)
}
