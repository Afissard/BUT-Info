
<?php
 function demo() {
  static $a = 0;
  echo ++$a.'<br>';
 }

 for ($i=0; $i<5; $i++) 
  demo();

?> 




 