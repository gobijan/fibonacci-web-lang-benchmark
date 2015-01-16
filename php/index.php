<?php 
	$n = $_GET['n'];
	
	function fib ($n = 1) {
		if ($n == 0) {
			return 0;
		} elseif ($n == 1) {
			return 1;
		} else {
			return fib($n - 1) + fib($n - 2);
		}
	}

	echo fib($n);
?>