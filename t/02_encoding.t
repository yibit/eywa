use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: hex
--- cmd

eywa encoding -type hex 313233343536375f6577654077737772255e58

--- args
--- out_like
1234567_ewe\@wswr\%\^X
--- err

=== TEST 2: hex -mode encoding
--- cmd

eywa encoding -type hex -mode encoding 313233343

--- args
--- out_like
333133323333333433
--- err
