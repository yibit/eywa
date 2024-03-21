use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: number to bits
--- cmd

eywa number -type bits -num 9

--- args
--- out_like
1001
--- err

=== TEST 2: count ones of number
--- cmd

eywa number -type ones -num 9

--- args
--- out_like
2
--- err

=== TEST 3: number to bits and count ones
--- cmd

eywa number -type default -num 9

--- args
--- out_like
1001:2
--- err

=== TEST 4: hex to int
--- cmd

eywa number -type int -n 0x1fff07

--- args
--- out_like
2096903
--- err
