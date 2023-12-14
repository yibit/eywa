use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: bits
--- cmd

eywa number -type bits -num 9

--- args
--- out_like
1001
--- err

=== TEST 2: bits
--- cmd

eywa number -type ones -num 9

--- args
--- out_like
2
--- err

=== TEST 3: default
--- cmd

eywa number -type default -num 9

--- args
--- out_like
1001:2
--- err
