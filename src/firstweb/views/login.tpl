<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Login</title>
	<script src="//cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
	<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
</head>
<body>
	<form class="form-inline" action="/trysql" method="post">
  <div class="form-group">
    <label for="exampleInputName2">Name</label>
    <input type="text" class="form-control" name="nickname" placeholder="nickname">
  </div>
  <div class="form-group">
    <label for="exampleInputEmail2">Email</label>
    <input type="text" class="form-control" name="password" placeholder="password">
  </div>
  <button type="submit" class="btn btn-default">Login In</button>
</form>
</body>
</html>