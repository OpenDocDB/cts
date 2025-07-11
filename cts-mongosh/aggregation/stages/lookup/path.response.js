response = {
"cursor": {
"firstBatch": [
{
"_id": "binary",
"_comment": "The base64 encoded string of bytes [42, 0, 13].",
"v": BinData(0, "KgAN"),
"composite_docs": [

]
},
{
"_id": "binary-empty",
"v": BinData(0, ""),
"composite_docs": [

]
},
{
"_id": "bool-false",
"v": false,
"composite_docs": [

]
},
{
"_id": "bool-true",
"v": true,
"composite_docs": [

]
},
{
"_id": "datetime",
"_comment": "The equivalent of 2021-11-01T10:18:42.123Z.",
"v": ISODate("2021-11-01T10:18:42.123Z"),
"composite_docs": [

]
},
{
"_id": "datetime-epoch",
"_comment": "Unix epoch time.",
"v": ISODate("1970-01-01T00:00:00Z"),
"composite_docs": [

]
},
{
"_id": "datetime-max",
"_comment": "The equivalent of 9999-12-31T23:59:59.999Z.",
"v": ISODate("9999-12-31T23:59:59.999Z"),
"composite_docs": [

]
},
{
"_id": "datetime-min",
"_comment": "The equivalent of 0000-01-01T00:00:00Z.",
"v": ISODate("0000-01-01T00:00:00Z"),
"composite_docs": [

]
},
{
"_id": "decimal128",
"v": Decimal128("3475653012423180288.4213"),
"composite_docs": [

]
},
{
"_id": "decimal128-infinity",
"v": Decimal128("8646911284551352320.0"),
"composite_docs": [

]
},
{
"_id": "decimal128-nan",
"v": Decimal128("8935141660703064064.0"),
"composite_docs": [

]
},
{
"_id": "decimal128-negative-infinity",
"v": Decimal128("17870283321406128128.0"),
"composite_docs": [

]
},
{
"_id": "decimal128-negative-zero",
"v": Decimal128("12699587999231377408.0"),
"composite_docs": [

]
},
{
"_id": "decimal128-whole",
"v": Decimal128("3476778912330022912.42"),
"composite_docs": [
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
}
]
},
{
"_id": "decimal128-zero",
"v": Decimal128("3476215962376601600.0"),
"composite_docs": [

]
},
{
"_id": "double",
"v": 42.13,
"composite_docs": [

]
},
{
"_id": "double-big",
"_comment": "Value x such as x+1 == x.",
"v": Double(2305843009213694000),
"composite_docs": [

]
},
{
"_id": "double-infinity",
"v": Double(+Infinity),
"composite_docs": [

]
},
{
"_id": "double-max",
"_comment": "The largest positive finite value representable as a double.",
"v": Double(179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000),
"composite_docs": [

]
},
{
"_id": "double-max-integer",
"_comment": "The largest integer value representable as a double.",
"v": Double(9007199254740991),
"composite_docs": [

]
},
{
"_id": "double-nan",
"v": Double(NaN),
"composite_docs": [

]
},
{
"_id": "double-negative-infinity",
"v": Double(-Infinity),
"composite_docs": [

]
},
{
"_id": "double-negative-zero",
"v": Double(-0),
"composite_docs": [

]
},
{
"_id": "double-smallest-non-zero",
"_comment": "The smallest positive non-zero value representable as a double.",
"v": 0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005,
"composite_docs": [

]
},
{
"_id": "double-whole",
"v": Double(42),
"composite_docs": [
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
}
]
},
{
"_id": "double-zero",
"v": Double(0),
"composite_docs": [

]
},
{
"_id": "int32",
"v": 42,
"composite_docs": [
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
}
]
},
{
"_id": "int32-max",
"v": 2147483647,
"composite_docs": [

]
},
{
"_id": "int32-min",
"v": -2147483648,
"composite_docs": [

]
},
{
"_id": "int32-zero",
"v": 0,
"composite_docs": [

]
},
{
"_id": "int64",
"v": Long(42),
"composite_docs": [
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
}
]
},
{
"_id": "int64-max",
"v": Long(9223372036854775807),
"composite_docs": [

]
},
{
"_id": "int64-min",
"v": Long(-9223372036854775808),
"composite_docs": [

]
},
{
"_id": "int64-zero",
"v": Long(0),
"composite_docs": [

]
},
{
"_id": "null",
"v": null,
"composite_docs": [
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
"_id": "document-empty",
"v": {

}
},
{
"_id": "document-null",
"v": {
"foo": null
}
}
]
},
{
"_id": "objectid",
"v": ObjectId("000102030405060708091011"),
"composite_docs": [

]
},
{
"_id": "objectid-empty",
"v": ObjectId("000000000000000000000000"),
"composite_docs": [

]
},
{
"_id": "string",
"v": "foo",
"composite_docs": [

]
},
{
"_id": "string-empty",
"v": "",
"composite_docs": [

]
},
{
"_id": "unset",
"composite_docs": [
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
"_id": "document-empty",
"v": {

}
},
{
"_id": "document-null",
"v": {
"foo": null
}
}
]
},
{
"_id": ObjectId("000000000000000000000000"),
"_comment": "_id is an empty ObjectID.",
"v": "objectid-empty",
"composite_docs": [

]
},
{
"_id": ObjectId("000102030405060708091011"),
"_comment": "_id is an ObjectID.",
"v": "objectid",
"composite_docs": [

]
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
