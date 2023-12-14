use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: json
--- cmd

eywa format -t json '{"array": ["foo","bar","baz"],"bool": false,"null": null,"num": 100,"obj": {"a":{"x": 1,"y": 2},"b": 2},"str": "foo"}'

--- args
--- out_like
\{
  "array": \[
    "foo",
    "bar",
    "baz"
  \],
  "bool": false,
  "null": null,
  "num": 100,
  "obj": \{
    "a": \{
      "x": 1,
      "y": 2
    \},
    "b": 2
  \},
  "str": "foo"
\}
--- err
