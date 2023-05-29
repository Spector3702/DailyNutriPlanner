document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的email文字
    var email = document.getElementById('email').value;


    //利用baseUrl加上email才是request Url
    const Url = 'http://127.0.0.1:8080/api/v1/persnInfo/?email=' + email;


    //利用get傳送email 給api 
    fetch(Url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
        }
    })
        .then(response => {
            if (response.ok) {
                // 登入成功
                return response.json().then(responseData => {
                    localStorage.setItem('responseData', JSON.stringify(responseData));
                    sessionStorage.setItem('isLogged', true);
                    window.location.href = "MainPage.html";
                });
            } else {
                // 登入失敗，跳轉到註冊頁面
                window.location.href = "RegisterPage.html";
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });

});

