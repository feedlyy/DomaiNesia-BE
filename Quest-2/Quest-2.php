<?php

$s = "";

$s = "charmander";

//get how many length in string that without repeating characters

//convert string to arr to loop
$arr = str_split($s);

//get list of each char
$ascii = [];

for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "Example 1:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";

//assign new value for another test
$s = "squirtle";
$arr = str_split($s);
$ascii = [];
for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "<br>Example 2:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";

//assign new value for another test
$s = "chaar";
$arr = str_split($s);
$ascii = [];
for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "<br>Example 3:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";

//assign new value for another test
$s = "chAar";
$arr = str_split($s);
$ascii = [];
for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "<br>Example 4:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";

//assign new value for another test
$s = "chaHr";
$arr = str_split($s);
$ascii = [];
for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "<br>Example 5:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";

//assign new value for another test
$s = "cha!!";
$arr = str_split($s);
$ascii = [];
for ($i = 0; $i < count($arr); $i++) {
    // keep original char
    $originalArrI = $arr[$i];

    /*check if the current char is :
    1. uppercase, then lowercase it
    2. lowercase, then uppercase it*/
    if (ord($arr[$i]) >= 65 && ord($arr[$i]) <= 90) {
        $arr[$i] = strtolower($arr[$i]);
    } elseif (ord($arr[$i]) >= 97 && ord($arr[$i]) <= 122) {
        $arr[$i] = strtoupper($arr[$i]);
    }

    /*check whether the changed value or the original char is
    already in the array*/
    if (in_array($arr[$i], $ascii) || in_array($originalArrI, $ascii)) {
        /*stop the loop because the current char is
        already in the list of $ascii*/
        break;
    } else {
        array_push($ascii, $originalArrI);
    }
}

echo "<br>Example 6:"."<br>";
echo "Input: ".$s."<br>";
echo "Output: ".count($ascii)."<br>";
echo "Explanation: The longest substring is ".implode("", $ascii).
    ", with the length of ".count($ascii)."<br>";
