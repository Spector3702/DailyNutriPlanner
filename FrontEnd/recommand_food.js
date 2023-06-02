document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的email文字
    var nutri = document.getElementById('select_nutri').value;
    var need = document.getElementById('select_amount').value;

    //利用baseUrl加上email才是request Url
    const Url = 'http://127.0.0.1:8080/api/v1/recNutri/food/?nutri=' + nutri + '&need=' + need;


    //利用get傳送email 給api 
    fetch(Url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
        }
    })
        .then(response => {
            return response.json();
        })
        .then(responseData => {
            console.log(responseData);
            var dataContainer = document.getElementById('dataContainer');
            var result = JSON.stringify(responseData.food);
            dataContainer.innerHTML = "";
            dataContainer.innerHTML = "你的隨機推薦食物為<br>"+result ;
        })
        .catch(error => {
            console.error('Error:', error);
        });

});

