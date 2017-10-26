function login_click(name,pwd) {
    $("#loginform").attr("action","user/login");
    $("#loginform").submit();
}