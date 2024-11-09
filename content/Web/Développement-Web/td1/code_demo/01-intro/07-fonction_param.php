<?php

//avec une boucle for
function moyenne1(Array $tab=array(10)):Float {
    $somme = 0;
    foreach ($tab as $elt) 
        $somme += $elt;
    if (count($tab)>0) 
        $somme /= count($tab);
    return $somme;        
}
//avec reduce et fonction annonyme
function moyenne2(Array $tab=array(10)):Float {
    $somme = array_reduce($tab,function($carry, $elt){
        return $carry + $elt;}, 0.0);
   
    if (count($tab)>0) 
        $somme /= count($tab);
    return $somme;        
}

echo moyenne1(array(10,12))."<br>".moyenne2(array(10,12))."<br>";
echo moyenne1(array()).'<br>'.moyenne2(array()).'<br>';
echo moyenne1().'<br>'.moyenne2().'<br>';
?>