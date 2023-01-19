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
?>
<div class="p-3">
<a href="php/logout.php" class="btn btn-outline-danger">Выйти!</a>
</div>
<div class="container">

</div>
<?php 
} else {
  ?>
<div class="container">
    <h2>Приветствую тебя, странник.</h2>
    <div class="container-register m-2">
        <h3>Регистрация!</h3>
        <form action="php/register.php" method="post">
            <div class="mb-3">
                <label for="exampleInputEmail1" class="form-label">Название компании</label>
                <input type="text" class="form-control" name="nameInput" aria-describedby="emailHelp">
              </div>
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">Адрес электронной почты</label>
              <input type="email" class="form-control" name="emailInput" aria-describedby="emailHelp">
            </div>
            <div class="mb-3">
              <label for="exampleInputPassword1" class="form-label">Пароль</label>
              <input type="password" class="form-control" name="passwordInput">
            </div>
            <div class="mb-3">
                <label for="exampleInputEmail1" class="form-label">Информация и комании</label>
                <input type="text" class="form-control" name="descriptiondInput" aria-describedby="emailHelp">
              </div>
              <div class="mb-3">
                <label for="exampleInputEmail1" class="form-label">Страна</label>
                <input type="text" class="form-control" name="countryInput" aria-describedby="emailHelp">
              </div>
              <div class="mb-3">
                <label for="exampleInputEmail1" class="form-label">Город</label>
                <input type="text" class="form-control" name="cityInput" aria-describedby="emailHelp">
              </div>
            <button type="submit" class="btn btn-primary">Регистрация</button>
          </form>
          <h5>Есть аккаунт? <a href="index.php">Авторизация</a></h5>

    </div>
    
   </div>
  <?php
}
?>
<!-- JavaScript Bundle with Popper -->
<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js" integrity="sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+jjXkk+Q2h455rYXK/7HAuoJl+0I4" crossorigin="anonymous"></script>
</body>
</html>