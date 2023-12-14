use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: crypto md5
--- cmd

eywa crypto -type md5 "pick"

--- args
--- out_like
27cf1e36682fb645e2f4943bbe54f7be
--- err
