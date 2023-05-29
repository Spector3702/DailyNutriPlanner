document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的email文字
    var food_bar = document.getElementById('food_bar').value;

    // 要處理中文
    const baseUrl = 'http://127.0.0.1:8080/api/v1/foodstuff/';
    const requestUrl = `${baseUrl}?name=${encodeURIComponent(food_bar)}`

    //利用get傳送email 給api 
    fetch(requestUrl, {
        method: 'GET',
        // mode: "no-cors",
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
        }
    })
        .then(response => {
            if (response.ok) {

                return response.json().then(responseData => {
                    console.log(responseData)
                });
            } else {

            }
        })
        .catch(error => {
            console.error('Error:', error);
        });

});

