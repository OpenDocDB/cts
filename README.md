# OpenDocDB Compatibility Test Suite

This repository contains a Compatibility Test Suite (CTS) for OpenDocDB-compatible databases
and a `opendocdb-cts` tool for managing CTS.

The `cts` directory contains JSON (Extended JSON v2 in Canonical Mode) files representing CTS:
fixtures (data sets), operations on those fixtures, and expected results.
Having them as data files (as opposed to programs or scripts in any programming language)
allows easy format conversions, integration into CI pipelines, and documentation.

`opendocdb-cts` directory contains the tool for generating parts of CTS,
running it against various targets and converting it to multiple formats.
