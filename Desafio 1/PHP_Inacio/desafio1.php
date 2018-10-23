<?php
$santa = "Santa";
$rem = "rem";
for ($cont=1;$cont<=200;$cont++):
    if ($cont % 6 == 0 and $cont % 5 == 0):
        print($cont." ".$santa.$rem. "\n");
    elseif ($cont % 5 == 0):
        print($cont." ".$santa. "\n");
    elseif ($cont % 6 == 0):
        print($cont." ".$rem. "\n");
    else:
        print($cont. "\n");
    endif;
endfor;
