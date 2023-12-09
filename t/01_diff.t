use lib 't/lib';
use strict;
use warnings;
use Test::Eywa skip_all => 'Skip: output is diff each time';

run_tests();

__DATA__

=== TEST 1: text
--- cmd

diff -t text

--- fileA
xcYcc

--- fileB
xcccc

--- args
--- out_like
xcYccc

--- err
