<?php

$position = $_POST['positionInput'];
$experience = $_POST['experienceInput'];
$description = $_POST['descriptionInput'];
$company_id = $_COOKIE['id'];


$url = '127.0.0.1:8080/vacancy/create';
$headers = ['Content-Type: application/json'];
$post_data = [
    'Position' => $position,
    'Experience' => $experience,
    'Description' => $description,
    'Company_id' => $company_id,
];
$data_json = json_encode($post_data);
$curl = curl_init();
curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
curl_setopt($curl, CURLOPT_VERBOSE, 1);
curl_setopt($curl, CURLOPT_POSTFIELDS, $data_json);
curl_setopt($curl, CURLOPT_URL, $url);
curl_setopt($curl, CURLOPT_POST, true);

$result = curl_exec($curl);

header("Location: ".$_SERVER['HTTP_REFERER']);
