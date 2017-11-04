<!DOCTYPE html>
<html>
<head>
  <meta charset=utf-8 />
  <title>Minesweeper</title>
  <link rel="stylesheet" type="text/css" media="screen" href="static/css/thunder/styles.css" />
  <!--[if IE]>
    <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
  <![endif]-->
</head>
<body>
  <div class="wrapper">
    <h1>Minesweeper</h1>
    <div id="minesweeper">
      <ul class="level-select">
        <li><button data-level="easy">Easy</button></li>
        <li><button data-level="mid">Medium</button></li>
        <li><button data-level="hard">Hard</button></li>
      </ul>
      <div class="menu">
        <ul>
          <li class="timer">000</li>
          <li class="reset"><button> :) </button></li>
          <li class="mine-count">010</li>
        </ul>
      </div>
      <div id="mine-grid"></div>
    </div>
  </div>
  <script src="static/js/jquery-2.2.4.min.js"></script>
  <script type="text/javascript" src="static/js/thunder/JSweeper.js"></script>
</body>
</html>