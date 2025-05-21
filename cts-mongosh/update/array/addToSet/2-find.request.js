db.runCommand({
"find": "composites",
"sort": {
"_id": 1
},
"multi": true
});
