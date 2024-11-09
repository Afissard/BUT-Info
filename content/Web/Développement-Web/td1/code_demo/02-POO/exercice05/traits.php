<?php


trait school {
    public function learn() {
        echo "I am learning PHP at WayToLearnX! \n";
    }
}
  trait company {
      public function work() {
          echo "I am working with PHP at WayToLearnX!";
      }
  }

  class Person {
      use school, company;
  }
  $person = new Person();
  $person->learn();
  $person->work();
?>
