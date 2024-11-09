<!-- 01-script.php -->
<pre>
Hors du php
<?php 
//echo permet un affichage
    echo "Dans du PHP ";
?>
Hors du php 
<?php
    echo "de retour en PHP";

    //
    $user = "jub";
    require "02-script.php";
    echo $user;

?>
</pre>