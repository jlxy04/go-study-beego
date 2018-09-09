<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <body>
    <form action="/customer/add" method="post">
        名字：<input type="text" name="name">
        年龄：<input type="number" name="age">
        性别：男<input type="radio" name="sex" value="男">女<input type="radio" name="sex" value="女">
        手机：<input type="number" name="mobile">
        <input type="submit" value="提交">
    </form>
  </body>
</html>
