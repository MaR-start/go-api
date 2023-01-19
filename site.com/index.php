<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <!-- CSS only -->
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-rbsA2VBKQhggwzxH7pPCaAqO46MgnOM80zW1RWuH61DGLwZJEdK2Kadq2F9CUG65" crossorigin="anonymous">
    <title>Главная</title>
</head>
<body>
<?php
if (isset($_COOKIE["auth_status"]) && $_COOKIE["auth_status"]== "ok") {
  $company_id = $_COOKIE["id"];
  $url = '127.0.0.1:8080/vacancy/find';
  $headers = ['Content-Type: application/json'];
  $post_data = [
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
  $arr = json_decode($result, true);
?>
<div class="p-3">
  <h4>Привет, <?= $_COOKIE["name"]; ?>!</h4>
  <button type="button" class="btn btn-outline-success" data-bs-toggle="modal" data-bs-target="#exampleModal">
  Создать
</button>
  <a href="php/logout.php" class="btn btn-outline-danger">Выйти!</a>
</div>

<div class="modal fade" id="exampleModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <h5 class="modal-title" id="exampleModalLabel">Создание вакансии</h5>
        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Закрыть"></button>
      </div>
      <div class="modal-body">
      <form action="php/create.php" method="post">
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">Должность</label>
              <input type="text" class="form-control" name="positionInput">
            </div>
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">Опыт работы</label>
              <input type="number" class="form-control" name="experienceInput" min="1">
            </div>
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">Описание должности</label>
              <input type="text" class="form-control" name="descriptionInput">
            </div>         
      </div>
      <div class="modal-footer">
        <button type="submit" class="btn btn-primary">Создать вакансию</button>
        <button type="button" class="btn btn-danger" data-bs-dismiss="modal">Закрыть</button>
        </form>
      </div>
    </div>
  </div>
</div>



<div class="container text-center">
  <h4>Доступные вакансии!</h4>
<div class="row">
<?php
foreach ($arr as $key) {
?>
<div class="col">
      <div class="card text-bg-light mb-3" style="max-width: 18rem;">
      <div class="card-header">№ <?= $key["id"] ?></div>
      <div class="card-body">
        <h5 class="card-title"><?= $key["position"] ?></h5>
        <p class="card-text"><?= $key["description"] ?></p>
        <p class="card-text"><small class="text-muted">Опыт работы <?php if ($key["experience"] == 0) { echo " не требуется"; } else if ($key["experience"] == 1) { echo " от " .$key['experience']. " года"; } else { echo " от " .$key['experience']. " лет"; }; ?> </small></p>
      </div>
    </div>
  </div>
<?php
}
?>
</div>
</div>
<?php 
} else {
  ?>
<div class="container">
    <h2>Приветствую тебя, странник.</h2>
    <div class="container-login m-2">
        <h3>Авторизация!</h3>
        <form action="php/login.php" method="post">
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">Адрес электронной почты</label>
              <input type="email" class="form-control" name="emailInput" aria-describedby="emailHelp">
              <div id="emailHelp" class="form-text">Мы никогда никому не передадим вашу электронную почту.</div>
            </div>
            <div class="mb-3">
              <label for="exampleInputPassword1" class="form-label">Пароль</label>
              <input type="password" class="form-control" name="passwordInput">
            </div>
            <button type="submit" class="btn btn-primary">Войти</button>
          </form>
          <h5>Нет аккаунта? <a href="register.php">Регистрация</a></h5>
    </div>    
   </div>
  <?php
}
?>
<!-- JavaScript Bundle with Popper -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+jjXkk+Q2h455rYXK/7HAuoJl+0I4" crossorigin="anonymous"></script>
</body>
</html>