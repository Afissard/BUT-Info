<?php
 $a = "toto";
 $b="a";

 echo $$b;

spl_autoload_register(function ($classe){
    echo $classe;die();
});

$obj1 = new un\Chien();

