$(function () {
    var bg_img = ['/static/img/login-bg1.jpg','/static/img/login-bg2.jpg','/static/img/login-bg3.jpg','/static/img/login-bg4.jpg','/static/img/login-bg5.jpg'];
    var index=0;

    $.backstretch(bg_img[index],{transitionDuration:800});

    setInterval(function(){
        index=(index>=bg_img.length-1)?0:index+1;
//            $.backstretch(bg_img[index]);
        $('#large-header').css('background-image','url('+bg_img[index]+')');
    },5000);
})

//回车事件
$(document).keydown(function(event){
    if(event.keyCode==13){
        subLogin();
    }
});

$('#login_btn').bind("click",function () {
    subLogin();
});

function subLogin() {
    $('#login_form').attr('action','/user/login');
    $('#login_form').submit();
}