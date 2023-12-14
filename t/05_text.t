use lib 't/lib';
use warnings;
use strict;
use Test::Eywa 'no_plan';

run_tests();

__DATA__

=== TEST 1: upper
--- cmd

eywa text -t upper "An Open guide to data structures and algorithms"

--- args
--- out_like
AN OPEN GUIDE TO DATA STRUCTURES AND ALGORITHMS
--- err

=== TEST 2: lower
--- cmd

eywa text -t lower "An Open guide to data structures and algorithms"

--- args
--- out_like
an open guide to data structures and algorithms
--- err

=== TEST 3: title
--- cmd

eywa text -t title "An Open guide to data structures and algorithms"

--- args
--- out_like
An Open Guide To Data Structures And Algorithms
--- err
