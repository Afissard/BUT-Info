
<?php
 //$a est global
 $a = "abc";
  
 function demo() {
   //Le scope de $a est ici local
   $a = 10; 
   //affiche $a si il est défini sinon le message non défini
    echo ($a??"non définie").'<br>';
 }

 
demo();

?>


 