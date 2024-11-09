
<?php
 //$a est global
 $a = "abc";
 //La variable magique $GLOBALS est un tableau associatif qui contiend 
 //les variables globales 
 echo $GLOBALS['a'].'<br>';
 
 function demo1() {
    //affiche $a si il est défini sinon le message non défini
    echo ($a??"non définie").'<br>';
 }

 function demo2() {
    //Le mot clef global permet d'accèder à une variable globale
    global $a;
    //affiche $a si il est défini sinon le message non défini
    echo ($a??"non définie").'<br>';
 }
demo1();
demo2();
?>


 