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
