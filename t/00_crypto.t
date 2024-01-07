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

=== TEST 2: crypto aes
--- cmd

eywa crypto -type aes 223ed888dcd216dcd40c47ff7cdaa7fd7eab65f4f0405350a43c5cad5b6b47b527c709edec29d7d6967518 2dd0b783290747ba62a63fc53591170d d2e6afe652c78f21 ISO78164

--- args
--- out_like
d24db664f2f5fc48e87fc07a225a37ff25aeacb027159f95b024ddbe153b9aaffa41af82e2717618d2de95b74c13163037098b7d30fdff18d0c1077f0fbf2284543c3754b7f981b60370658da102da9965075638fa767915c7106a6e299efd5a
--- err

=== TEST 3: crypto des
--- cmd

eywa crypto -type des helloworld world123

--- args
--- out_like
4c8b3058381e076eedbaf2c1291654ea
--- err
