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