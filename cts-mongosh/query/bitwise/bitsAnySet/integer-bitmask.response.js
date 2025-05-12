response = {
"cursor": {
"firstBatch": [
{
"_id": "binary",
"_comment": "The base64 encoded string of bytes [42, 0, 13].",
"v": BinData(128, "KgAN")
},
{
"_id": "decimal128-whole",
"v": Decimal128("3476778912330022912.42")
},
{
"_id": "double-max-integer",
"_comment": "The largest integer value representable as a double.",
"v": Double(9007199254740991)
},
{
"_id": "double-whole",
"v": Double(42)
},
{
"_id": "int32",
"v": 42
},
{
"_id": "int32-max",
"v": 2147483647
},
{
"_id": "int64",
"v": Long(42)
},
{
"_id": "int64-max",
"v": Long(9223372036854775807)
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
