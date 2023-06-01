document.addEventListener('DOMContentLoaded', function () {
    var isLoggedIn = sessionStorage.getItem('isLogged');

    if (isLoggedIn) {
        console.log(responseData);
    } else {
        window.location.href = "LoginPage.html";
    }
});