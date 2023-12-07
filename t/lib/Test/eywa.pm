package Test::Eywa;

use Test::Base -Base;
use File::Temp qw/ :POSIX /;
use IPC::Run ();


our @EXPORT = qw( run_tests );

sub run_tests () {
    for my $block (Test::Base::blocks()) {
        run_test($block);
    }
}

sub bail_out (@) {
    Test::More::BAIL_OUT(@_);
}

sub parse_cmd ($) {
    my $cmd = shift;
    my @cmd;

    while (1) {
        if ($cmd =~ /\G\s*"(.*?)"/gmsc) {
            push @cmd, $1;

        } elsif ($cmd =~ /\G\s*'(.*?)'/gmsc) {
            push @cmd, $1;

        } elsif ($cmd =~ /\G\s*(\S+)/gmsc) {
            push @cmd, $1;

        } else {
            last;
        }
    }

    return @cmd;
}

sub run_test ($) {
    my $block = shift;
    my $name = $block->name;

    my $timeout = $block->timeout() || 10;
    my $opts = $block->opts;
    my $args = $block->args;

    my $cmd = "bin/eywa";

    my $ssfile;
    if (defined $opts) {
        $cmd .= " $opts";
    }

    if (defined $block->cmd) {
        $cmd .= " " . $block->cmd;
    }

    if (defined $args) {
    	$cmd .= " $args";
    }

    #warn "CMD: $cmd\n";

    my @cmd = parse_cmd($cmd);

    my ($out, $err);

    eval {
        IPC::Run::run(\@cmd, \undef, \$out, \$err,
                      IPC::Run::timeout($timeout));
    };

    if ($@) {

        if ($@ =~ /timeout/) {

            if (!defined $block->expect_timeout) {
                fail("$name: meta process timeout");
            }

        } else {
            fail("$name: failed to run command [$cmd]: $@");
        }
    }

    my $ret = ($? >> 8);

    if (defined $ssfile) {
        unlink $ssfile;
    }

    if (defined $block->out) {
        is $out, $block->out, "$name - stdout eq okay";
    }

    my $regex = $block->out_like;
    if (defined $regex) {
        if (!ref $regex) {
            $regex = qr/$regex/ms;
        }
        like $out, $regex, "$name - stdout like okay";
    }

#    if (defined $block->err) {
#        is $err, $block->err, "$name - stderr eq okay";
#    }

    $regex = $block->err_like;
    if (defined $regex) {
    	if (!ref $regex) {
    	    $regex = qr/$regex/ms;
    	}
    	like $err, $regex, "$name - stderr like okay";
    }

}

1;
