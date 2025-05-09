response = {
"cursor": {
"firstBatch": [
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
}
],
"id": Long(0),
"ns": "db.composites"
},
"ok": Double(1)
}
