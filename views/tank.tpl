<!DOCTYPE html>
<html lang="zh" class="no-js demo-1">
	<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"> 
		<meta name="viewport" content="width=device-width, initial-scale=1.0"> 
		<script src="static/js/tank/jquery.min.js"></script>
		<script src="static/js/tank/Helper.js"></script>
		<script src="static/js/tank/keyboard.js"></script>
		<script src="static/js/tank/const.js"></script>
		<script src="static/js/tank/level.js"></script>
		<script src="static/js/tank/crackAnimation.js"></script>
		<script src="static/js/tank/prop.js"></script>
		<script src="static/js/tank/bullet.js"></script>
		<script src="static/js/tank/tank.js"></script>
		<script src="static/js/tank/num.js"></script>
		<script src="static/js/tank/menu.js"></script>
		<script src="static/js/tank/map.js"></script>
		<script src="static/js/tank/Collision.js"></script>
		<script src="static/js/tank/stage.js"></script>
		<script src="static/js/tank/main.js"></script>
		<style type="text/css">
			#canvasDiv canvas{
				position:absolute;
			}
		</style>
	</head>
	<body>
		<div class="container">
			<head><h3>操作说明：玩家1：wasd上左下右，space射击；玩家2：方向键，enter射击。n下一关，p上一关。</h3></head>
			<div class="main clearfix">
				<div id="canvasDiv" >
					<canvas id="wallCanvas" ></canvas> 
					<canvas id="tankCanvas" ></canvas>
					<canvas id="grassCanvas" ></canvas>
					<canvas id="overCanvas" ></canvas> 
					<canvas id="stageCanvas" ></canvas>
				</div>
			</div>
			
		</div><!-- /container -->
	</body>
</html>