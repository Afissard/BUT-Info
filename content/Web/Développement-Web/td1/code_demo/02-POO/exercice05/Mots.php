<?php

class Mots
{
    private $mots = [];

    public function getItems()
    {
        return $this->mots;
    }

    public function addItem($mot)
    {
        $this->mots[] = $mot;
    }

    
}
