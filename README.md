# OpenDocDB Compatibility Test Suite

This repository contains a Compatibility Test Suite (CTS) for OpenDocDB-compatible databases
and a `opendocdb-cts` tool for managing CTS.

See [this blog post](https://opendocdb.org/steps-to-consider-when-standardizing-a-json-like-query-language) for some background information.

## CTS

The `cts` directory contains files representing the Compatibility Test Suite.

The building block of CTS is the *test case*: a pair of the request to be send to the OpenDocDB-compatible database and the expected response.
A logical group of test cases is called a *test suite*.
Requests in the test suite use pre-populated data from *fixtures*.

Example fixture:

```json
[
  {"_id": "double-zero",          "v": {"$numberDouble": "0.0"}},
  {"_id": "double-negative-zero", "v": {"$numberDouble": "-0.0"}},
  // ...
]
```

Example test case:

```json
{
  "request": {
    "find": "values",
    "filter": {
      "v": {"$eq": {"$numberDouble": "0.0"}}
    },
    "sort": {
      "_id": {"$numberInt": "1"}
    },
    "$db": "{{.Database}}"
  },
  "response": {
    "cursor": {
      "firstBatch": [
        {"_id": "double-negative-zero", "v": {"$numberDouble": "-0.0"}},
        {"_id": "double-zero",          "v": {"$numberDouble": "0.0"}},
        // ...
      ],
      "id": {"$numberLong": "0"},
      "ns": "{{.Database}}.values"
    },
    "ok": {"$numberDouble": "1.0"}
  }
}
```

Test cases use the following data types:
* document (object);
* array;
* double;
* string (UTF-8 string);
* binary (binary data);
* ObjectId;
* boolean;
* date (UTC datetime);
* null;
* regex;
* int32;
* timestamp;
* int64;
* decimal128.

Document field names can't be empty.
Duplicate field names are not allowed.
The order of fields is currently mostly ignored.
All that might change in the future.

Double and decimal128 values include things like negative zero, infinity, NaN, etc.

Test cases, test suites, and fixtures are stored in Extended JSON v2 in Canonical Mode. Having them as data files (as opposed to programs or scripts in any programming language) allows easy format conversions, integration into CI pipelines, and documentation. The use of Extended JSON has three limitations that might be problematic:
1. Most JSON parsers do not preserve the order of fields in the document. Currently, CTS does not rely on the order of fields (see above), but that might change.
2. The use `$` for extended data types makes it impossible to test documents with fields with `$` in them.
3. Some values (like various forms of NaN) have the same representation.

## opendocdb-cts

`opendocdb-cts` directory contains the tool for running CTS against various OpenDocDB-compatible databases
using the wire protocol.
It also could convert fixtures and test cases to `mongosh` JavaScript syntax for documentation purposes.

## Roadmap

* [Generate test cases](https://github.com/OpenDocDB/cts/issues/5).
* Run CTS automatically against MongoDB, FerretDB, Azure Cosmos DB, and Amazon DocumentDB.
