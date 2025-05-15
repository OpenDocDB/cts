db.runCommand({
"find": "values",
"filter": {
"$jsonSchema": {
"bsonType": "object",
"properties": {
"v": {
"bsonType": "double"
}
}
}
},
"sort": {
"_id": 1
}
});
