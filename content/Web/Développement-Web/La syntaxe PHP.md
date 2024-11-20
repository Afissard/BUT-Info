# Variables, Classes, Fonctions
```php
// ouvrir la balise php
<?php

// déclarer une variable
$varname = varvalue

// déclarer une classe
class ClassName {

}

// déclarer une classe finale héritante (il est possible de déclarer une classe seulement comme finale ou héritante)
final class ClassName extends ClassToExtend {

}

// déclarer une variable dans une classe
public $var1;
protected $var2;
private $var3;

// déclarer une variable statique
public static $staticvar;

// modifier le constructeur
public function __construct(Type1 $var1, Type2 $var2) {

}

/*
Note : Une variable peut être dans un type au choix, exemple : 
Float | Int | String $var sera soit un Float, soit un Int, soit un String.
Elle peut également ne pas déclarer de type du tout.
*/

// modifier une variable dans sa classe
$this->var1 = newValue;

// modifier une variable statique dans sa classe
self::$staticvar = newValue;

// récupérer / modifier une variable statique d'une autre classe
NomDeLaClasse::staticvar

// lancer une exception
throw new CustomException('Custom Exception Here');

// afin d'override une fonction, il est plus robuste de déclarer :
#[\Override]
public function funToOverride() {}

// déclarer une enum class
enum ThisEnumClass : TypeOfElements {
	case one = 1; // par exemple
	case two = 2;
	case three = 3;
}

```

# Arrays
```php
// initialiser un Array
$list = [];

/*
Note : Une liste fonctionne par position ou par clé, en fonction de la manière dont on y ajoute des valeurs
*/

// vérifier si une clé existe déjà
array_key_exists($list);

// récupérer la longueur d'un Array
sizeOf($list);

// récupérer une liste des clés
$keys = array_keys($list);

// récupérer une valeur associée à une clé
$var = $list[$key];

// associer une valeur à une clé
$list[$key] = $var;

// créer un array avec des valeurs par clé
$list = array(
	"val1" => 1,
	"val2" => 2,
);

// créer un array avec des valeurs
$list = array(12, 40, 1, 34, 20)

/*
Note : Un array par position est équivalent à un array par clé dont les clés sont 0, 1, 2..
Il est donc possible de déclarer un array(17, 2, 8=>3) et d'avoir donc:
Array(
	[0] -> 17
	[1] -> 2
	[8] -> 3
)
*/
```

# Traits, Use
```php
/* un trait permet de définir des fonction que les classes peuvent utiliser via le keyword "use" */

// définir un trait
trait mytrait {
	public function fun1() {
		return "this is fun 1";
	}
	
	private function fun2() {
		return "this is fun 2";
	}
}

// "hériter" d'un trait
class MyClass {
	use mytrait;
}
```

# Dépendances
```php
// ajouter une dépendance
require "MyPath/File.php";

// créer une fonction présente dans cette dépendance
$object = new MyPath\Class();

// définir une classe venant d'une dépendance comme "utilisée"
use MyPath\Class;
$object = new Class();

// définir une classe comme "utilisée" sous un autre nom
use MyPath\Class as PathClass;
$object = new PathClass();
```

# Autoload
```php
/*
Note : Un autoload permet de gérer automatiquement les dépendances en faisant une opération spécifique définie dans cette fonction dès qu'il détecte qu'une dépendance n'est pas satisfaite.
*/

// définir une fonction autoload avec un pattern spécifique, ici chercher dans le dossier class le fichier nomclasse.php
spl_autoload_register(function($class) {
    require 'class/' . $class . '.php';
});

```