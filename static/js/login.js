$('#login_btn').bind("click",function () {
    $('#login_form').attr('action','/user/login');
    $('#login_form').submit();
});