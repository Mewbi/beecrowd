<?php
function euclideanDivision($a, $b) {
    $limit = $b < 0 ? -$b : $b;
    $r = 0;
    $q = 0;

    for ($r = 0; $r < $limit; $r++) {
        if (($a - $r) % $b == 0) {
            $q = ($a - $r) / $b;
            break;
        }
    }

    return array($q, $r);
}

list($a, $b) = explode(" ", fgets(STDIN));
$a = intval($a);
$b = intval($b);

list($q, $r) = euclideanDivision($a, $b);
echo "$q $r\n";
?>
