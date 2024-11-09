<?php
    require "un/A.php";
    require "deux/A.php";


    echo (new un\A())->a;  //un
    echo (new deux\A())->a;  //deux

    use un\A;
    echo (new A())->a;  //un

    use deux\A as AA;
    echo (new AA())->a;  //deux

?>
