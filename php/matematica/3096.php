<?php

function kamenetsky($n) {
    if ($n < 0) {
        return 0;
    }

    if ($n <= 1) {
        return 1;
    }

    $x = (($n * log10($n/M_E)) +
          (log10(2*M_PI*$n) / 2.0));

    return (int)(floor($x) + 1);
}

$input = fgets(STDIN);
$n = (float)$input;
echo kamenetsky($n) . "\n";

?>
