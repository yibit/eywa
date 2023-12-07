use lib 't/lib';
use strict;
use warnings;
use Test::Eywa skip_all => 'Skip: output is diff each time';

run_tests();

__DATA__

=== TEST 1: text
--- cmd

diff -t text a.txt b.txt

--- fileA: a.txt
xcYcc

--- filleB: b.txt
xcccc

--- args
--- out_like
xcYccc
--- err
