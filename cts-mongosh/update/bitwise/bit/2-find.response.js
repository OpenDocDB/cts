response = {
"cursor": {
"firstBatch": [
{
"_id": "binary",
"v": BinData(0, "KgAN")
},
{
"_id": "binary-empty",
"v": BinData(0, "")
},
{
"_id": "bool-false",
"v": false
},
{
"_id": "bool-true",
"v": true
},
{
"_id": "datetime",
"v": ISODate("2021-11-01T10:18:42.123Z")
},
{
"_id": "datetime-epoch",
"v": ISODate("1970-01-01T00:00:00Z")
},
{
"_id": "datetime-max",
"v": ISODate("9999-12-31T23:59:59.999Z")
},
{
"_id": "datetime-min",
"v": ISODate("0000-01-01T00:00:00Z")
},
{
"_id": "decimal128",
"v": Decimal128("3475653012423180288.4213")
},
{
"_id": "decimal128-infinity",
"v": Decimal128("8646911284551352320.0")
},
{
"_id": "decimal128-nan",
"v": Decimal128("8935141660703064064.0")
},
{
"_id": "decimal128-negative-infinity",
"v": Decimal128("17870283321406128128.0")
},
{
"_id": "decimal128-negative-zero",
"v": Decimal128("12699587999231377408.0")
},
{
"_id": "decimal128-whole",
"v": Decimal128("3476778912330022912.42")
},
{
"_id": "decimal128-zero",
"v": Decimal128("3476215962376601600.0")
},
{
"_id": "double",
"v": 42.13
},
{
"_id": "double-big",
"v": Double(2305843009213694000)
},
{
"_id": "double-infinity",
"v": Double(+Infinity)
},
{
"_id": "double-max",
"v": Double(179769313486231570000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000)
},
{
"_id": "double-max-integer",
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
"_id": "int32",
"v": 46
},
{
"_id": "int32-max",
"v": 2147483647
},
{
"_id": "int32-min",
"v": -2147483644
},
{
"_id": "int32-zero",
"v": 4
},
{
"_id": "int64",
"v": Long(46)
},
{
"_id": "int64-max",
"v": Long(9223372036854775807)
},
{
"_id": "int64-min",
"v": Long(-9223372036854775804)
},
{
"_id": "int64-zero",
"v": Long(4)
},
{
"_id": "null",
"v": null
},
{
"_id": "objectid",
"v": ObjectId("000102030405060708091011")
},
{
"_id": "objectid-empty",
"v": ObjectId("000000000000000000000000")
},
{
"_id": "string",
"v": "foo"
},
{
"_id": "string-empty",
"v": ""
},
{
"_id": "unset"
},
{
"_id": ObjectId("000000000000000000000000"),
"v": "objectid-empty"
},
{
"_id": ObjectId("000102030405060708091011"),
"v": "objectid"
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
