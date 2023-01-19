<?php

setcookie('auth_status', '', time()-3600, '/');
setcookie('name', '', time()-3600, '/');
setcookie('id', '', time()-3600, '/');
header("Location: ".$_SERVER['HTTP_REFERER']);