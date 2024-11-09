<pre>
<?php 
    //la variable $a a t elle été declarée 
    //et est elle differente de null
    //var_dump affichage pour debbuggage et isset pour tester 
    var_dump(isset($a));

    //declaration de $a et affectation de 10
    // $a est un int    
    $a = 10;
    var_dump($a);
    
    //$a est naitenant de type null et vaut null
    $a = null;
    var_dump($a);

    //$a est maintenant un tableau association 
    $a=array();
    var_dump($a);

    //ajout d'un élément au tableau
    $a['clef']='une valeur';
    var_dump($a);
    //ajout d'un élément au tableau
    $a[0]=10;
    var_dump($a);

    //$a n'est plus définie
    unset($a);
    var_dump(isset($a));
?>
</pre>