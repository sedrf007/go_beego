<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Welcome!</title>
	<link rel="stylesheet" href="//cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
</head>
<body>
	<form class="form-horizontal" action="/say" method="post">
  <div class="form-group">
    <label for="username" class="col-sm-2 control-label">输入你的用户名</label>
    <div class="col-sm-8">
      <input type="text" class="form-control" name="username" placeholder="">
    </div>
  </div>
  <div class="form-group">
    <div class="col-sm-offset-2 col-sm-10">
      <button type="submit" class="btn btn-default">进入聊天室</button>
    </div>
  </div>
</form>
	
</body>
</html>