<?php
require "Mots.php";
$collection = new Mots();
$collection->addItem("First");
$collection->addItem("Second");
$collection->addItem("Third");
$collection->addItem("Fourth");


foreach ($collection->getIterator() as $item) {
    echo $item . "\n";
}
?>
