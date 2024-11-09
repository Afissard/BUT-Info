<?php require "Etudiant.php"; ?>
<!doctype html>
<html lang="fr">
<head>
    <meta charset="utf-8">
    <title>Test Etudiant </title>
</head>
<body>
<pre>
<?php
    $etudiant1 = new Etudiant("E001",10);
    var_dump($etudiant1);
    $etudiant2 = new Etudiant("E002",moyenne: 20);
    var_dump($etudiant2);

    echo $etudiant1->getNo()." : ".$etudiant1->getMoyenne();
    echo '<br>';
    echo "{$etudiant2->getNo()} : {$etudiant2->getMoyenne()}";

    echo '<br>';
    echo $etudiant1->moyenne;
    echo '<br>';
    echo $etudiant1->no;
    echo '<br>';
    echo $etudiant2;
    echo '<br>';
    $etudiant3 = new Etudiant("E002",moyenne: 20);
    var_dump($etudiant3 === $etudiant2);
    var_dump($etudiant3 == $etudiant2);
    $etudiant4 = $etudiant3;
    var_dump($etudiant4 === $etudiant3);
    var_dump($etudiant4 == $etudiant3);
    ?>
</pre>
</body>
