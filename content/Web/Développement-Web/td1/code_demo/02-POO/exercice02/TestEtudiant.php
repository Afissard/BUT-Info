<?php require "Etudiant.php"; ?>
<!doctype html>
<html lang="fr">
<head>
    <meta charset="utf-8">
    <title>Test Etudiant 2 </title>
</head>
<body>
<pre>
<?php
    $etudiant1 = new Etudiant("E001",10);
    echo $etudiant1;
    $etudiant2 = new Etudiant("E002",moyenne: 20);
    echo '<br>';
    echo $etudiant2;
    echo '<br>';
    var_dump(Etudiant::$etatMental);
    Etudiant::$etatMental = EtatMental::Joyeux;
    echo '<br>';
    echo $etudiant2;
    echo '<br>';
    echo Etudiant::getNbEtudiant();
    ?>
</pre>
</body>
