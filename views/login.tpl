<!DOCTYPE html>
<html lang="en" class="no-js">
<head>
  <meta charset="UTF-8" />
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>登录</title>
  <link rel="Shortcut Icon" type="image/x-icon" href="/static/img/windows.ico">
  <link rel="stylesheet" type="text/css" href="/static/css/normalize.css" />
  <link rel="stylesheet" type="text/css" href="/static/css/demo.css" />
  <!--必要样式-->
  <link rel="stylesheet" type="text/css" href="/static/css/component.css" />

</head>
<body>
<div class="container demo-1">
  <div class="content">
    <div id="large-header" class="large-header" style="background-image: url('/static/img/login-bg1.jpg');">
      <a href="qq" class="chat">toChat</a>
      <!--
      <canvas id="demo-canvas"></canvas>
      -->
      <div class="logo_box">
        <h3>欢迎你</h3>
        <form action="#" id="login_form" method="post">
          <div class="input_outer">
            <span class="u_user"></span>
            <input name="username" class="text" style="color: #FFFFFF !important" type="text" placeholder="请输入账户" value="admin">
          </div>
          <div class="input_outer">
            <span class="us_uer"></span>
            <input name="password" class="text" style="color: #FFFFFF !important; position:absolute; z-index:100;"value="admin" type="password" placeholder="请输入密码">
          </div>
          <div class="mb2"><a id="login_btn" class="act-but submit" href="javascript:" style="color: #FFFFFF">登录</a></div>
        </form>
      </div>
    </div>
  </div>
</div><!-- /container -->
</div>
</body>
<!--
<script src="static/js/TweenLite.min.js"></script>
<script src="static/js/EasePack.min.js"></script>
<script src="static/js/rAF.js"></script>
-->
<script src="/static/js/jquery-1.9.1.min.js"></script>
<script src="/static/js/jquery.backstretch.min.js"></script>
<script src="/static/js/login.js"></script>
<script src="/static/js/login-sub.js"></script>
<script type="text/javascript">
    {{if .status}}
    alert({{.result}})
    {{end}}
</script>
</html>