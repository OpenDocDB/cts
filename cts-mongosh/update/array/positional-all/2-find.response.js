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
100,
100,
100,
100,
100,
100,
100,
100,
100,
100
]
},
{
"_id": "array-composite-reverse",
"v": [
100,
100,
100,
100,
100,
100,
100,
100,
100,
100
]
},
{
"_id": "array-document-empty",
"v": [
100
]
},
{
"_id": "array-documents",
"_comment": "An array of documents, add more documents if necessary.",
"v": [
100,
100,
100
]
},
{
"_id": "array-double-big",
"_comment": "The array contains the value x such that x+1 == x.",
"v": [
100
]
},
{
"_id": "array-double-desc",
"v": [
100,
100,
100
]
},
{
"_id": "array-double-duplicate",
"_comment": "The first and the second elements are the same value.",
"v": [
100,
100,
100
]
},
{
"_id": "array-double-max-integer",
"_comment": "The array contains the largest integer value representable as a double.",
"v": [
100
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
100
]
},
{
"_id": "array-null",
"v": [
100
]
},
{
"_id": "array-numbers-asc",
"v": [
100,
100,
100,
100
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
