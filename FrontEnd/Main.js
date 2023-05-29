document.addEventListener('DOMContentLoaded', function () {
    var isLoggedIn = sessionStorage.getItem('isLoggedIn');

    if (isLoggedIn) {
        console.log(responseData);
    } else {
        window.location.href = "LoginPage.html";
    }
});