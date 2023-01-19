<?php

$email = $_POST['emailInput'];
$password = $_POST['passwordInput'];


$url = '127.0.0.1:8080/company/login';

$headers = ['Content-Type: application/json'];

$post_data = [
    'Email' => $email,
    'Password' => $password,
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

$arr = json_decode($result, true);
if ($result !== "") {
    setcookie('auth_status', 'ok', time()+3600, '/');
    setcookie('name', $arr["name"], time()+3600, '/');
    setcookie('id', $arr["id"], time()+3600, '/');
    header("Location: ".$_SERVER['HTTP_REFERER']);

} else {
    echo $result;
    echo "<a href='/'>Назад!</a>";
}