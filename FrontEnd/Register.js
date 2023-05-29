document.getElementById('Registerform').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的文字
    var email = document.getElementById('email').value;
    var gender = document.getElementById('gender').value;
    var age = document.getElementById('age').value;
    var height = document.getElementById('height').value;
    var weight = document.getElementById('weight').value;
    var workload = document.getElementById('workload').value;

    //建一個變數來儲存剛剛的那些資料
    var formData = new URLSearchParams();
    formData.append('email', email);
    formData.append('gender', gender);
    formData.append('age', age);
    formData.append('height', height);
    formData.append('weight', weight);
    formData.append('workload', workload);


    const Url = 'http://127.0.0.1:8080/api/v1/persnInfo/';

    //利用fetch 的post傳送formData給api，讓他們新增
    fetch(Url, {
        method: 'POST',
        headers: new Headers({
            'Content-Type': 'application/x-www-form-urlencoded'
        }),
        body: formData
    })
        .then(response => {
            if (!response.ok) {
                window.location.href = "LoginPage.html";
            }
            return response.json();
        })
        .then(responseData => {
            console.log(responseData);
            // 註冊成功後就用本地端儲存responseData，然後就跳轉到MainPage
            localStorage.setItem('responseData', JSON.stringify(responseData));
            window.location.href = "MainPage.html";
        })
        .catch(error => {
            console.error('Error:', error);
        });

});
