document.getElementById('createform').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的文字
    var food = document.getElementById('food_bar').value;
    var data = localStorage.getItem('responseData');
    var parsedData = JSON.parse(data);
    var email = parsedData['Email'];

    //建一個變數來儲存剛剛的那些資料
    var formData = new URLSearchParams();
    formData.append('email', email);
    formData.append('name', food);

    const Url = 'http://127.0.0.1:8080/api/v1/record/create/';

    //利用fetch 的post傳送formData給api，讓他們新增
    fetch(Url, {
        method: 'POST',
        headers: new Headers({
            'Content-Type': 'application/x-www-form-urlencoded'
        }),
        body: formData
    })
        .then(response => {
            return response.json();
        })
        .then(responseData => {
            console.log(responseData);
            var dataContainer = document.getElementById('dataContainer');
            dataContainer.innerHTML = "";
            dataContainer.innerHTML = "新增成功!";
            
        })
        .catch(error => {
            console.error('Error:', error);
        });

});
