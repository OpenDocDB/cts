response = {
"cursor": {
"firstBatch": [
{
"_id": "binary",
"_comment": "The base64 encoded string of bytes [42, 0, 13].",
"v": BinData(0, "KgAN"),
"foo": "bar"
},
{
"_id": "binary-empty",
"v": BinData(0, ""),
"foo": "bar"
},
{
"_id": "bool-false",
"v": false,
"foo": "bar"
},
{
"_id": "bool-true",
"v": true,
"foo": "bar"
},
{
"_id": "datetime",
"_comment": "The equivalent of 2021-11-01T10:18:42.123Z.",
"v": ISODate("2021-11-01T10:18:42.123Z"),
"foo": "bar"
},
{
"_id": "datetime-epoch",
"_comment": "Unix epoch time.",
"v": ISODate("1970-01-01T00:00:00Z"),
"foo": "bar"
},
{
"_id": "datetime-max",
"_comment": "The equivalent of 9999-12-31T23:59:59.999Z.",
"v": ISODate("9999-12-31T23:59:59.999Z"),
"foo": "bar"
},
{
"_id": "datetime-min",
"_comment": "The equivalent of 0000-01-01T00:00:00Z.",
"v": ISODate("0000-01-01T00:00:00Z"),
"foo": "bar"
},
{
"_id": "decimal128",
"v": Decimal128("3475653012423180288.4213"),
"foo": "bar"
},
{
"_id": "decimal128-infinity",
"v": Decimal128("8646911284551352320.0"),
"foo": "bar"
},
{
"_id": "decimal128-nan",
"v": Decimal128("8935141660703064064.0"),
"foo": "bar"
},
{
"_id": "decimal128-negative-infinity",
"v": Decimal128("17870283321406128128.0"),
"foo": "bar"
},
{
"_id": "decimal128-negative-zero",
"v": Decimal128("12699587999231377408.0"),
"foo": "bar"
},
{
"_id": "decimal128-whole",
"v": Decimal128("3476778912330022912.42"),
"foo": "bar"
},
{
"_id": "decimal128-zero",
"v": Decimal128("3476215962376601600.0"),
"foo": "bar"
},
{
"_id": "double",
"v": 42.13,
"foo": "bar"
},
{
"_id": "double-big",
"_comment": "Value x such as x+1 == x.",
"v": Double(2305843009213694000),
"foo": "bar"
},
{
"_id": "double-infinity",
"v": Double(+Infinity),
"foo": "bar"
},
{
"_id": "double-max",
"_comment": "The largest positive finite value representable as a double.",
"v": Double(179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000),
"foo": "bar"
},
{
"_id": "double-max-integer",
"_comment": "The largest integer value representable as a double.",
"v": Double(9007199254740991),
"foo": "bar"
},
{
"_id": "double-nan",
"v": Double(NaN),
"foo": "bar"
},
{
"_id": "double-negative-infinity",
"v": Double(-Infinity),
"foo": "bar"
},
{
"_id": "double-negative-zero",
"v": Double(-0),
"foo": "bar"
},
{
"_id": "double-smallest-non-zero",
"_comment": "The smallest positive non-zero value representable as a double.",
"v": 0.000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000005,
"foo": "bar"
},
{
"_id": "double-whole",
"v": Double(42),
"foo": "bar"
},
{
"_id": "double-zero",
"v": Double(0),
"foo": "bar"
},
{
"_id": "int32",
"v": 42,
"foo": "bar"
},
{
"_id": "int32-max",
"v": 2147483647,
"foo": "bar"
},
{
"_id": "int32-min",
"v": -2147483648,
"foo": "bar"
},
{
"_id": "int32-zero",
"v": 0,
"foo": "bar"
},
{
"_id": "int64",
"v": Long(42),
"foo": "bar"
},
{
"_id": "int64-max",
"v": Long(9223372036854775807),
"foo": "bar"
},
{
"_id": "int64-min",
"v": Long(-9223372036854775808),
"foo": "bar"
},
{
"_id": "int64-zero",
"v": Long(0),
"foo": "bar"
},
{
"_id": "null",
"v": null,
"foo": "bar"
},
{
"_id": "objectid",
"v": ObjectId("000102030405060708091011"),
"foo": "bar"
},
{
"_id": "objectid-empty",
"v": ObjectId("000000000000000000000000"),
"foo": "bar"
},
{
"_id": "string",
"v": "foo",
"foo": "bar"
},
{
"_id": "string-empty",
"v": "",
"foo": "bar"
},
{
"_id": "unset",
"foo": "bar"
},
{
"_id": ObjectId("000000000000000000000000"),
"_comment": "_id is an empty ObjectID.",
"v": "objectid-empty",
"foo": "bar"
},
{
"_id": ObjectId("000102030405060708091011"),
"_comment": "_id is an ObjectID.",
"v": "objectid",
"foo": "bar"
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
