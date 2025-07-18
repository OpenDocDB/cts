response = {
"cursor": {
"firstBatch": [
{
"_id": "binary",
"_comment": "The base64 encoded string of bytes [42, 0, 13].",
"v": BinData(0, "KgAN")
},
{
"_id": "binary-empty",
"v": BinData(0, "")
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
"_id": "double-big",
"_comment": "Value x such as x+1 == x.",
"v": Double(2305843009213694000)
},
{
"_id": "double-negative-zero",
"v": Double(-0)
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
"v": 42
},
{
"_id": "int32-min",
"v": -2147483648
},
{
"_id": "int32-zero",
"v": 0
},
{
"_id": "int64",
"v": Long(42)
},
{
"_id": "int64-min",
"v": Long(-9223372036854775808)
},
{
"_id": "int64-zero",
"v": Long(0)
}
],
"id": Long(0),
"ns": "db.values"
},
"ok": Double(1)
}
