document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的email文字
    const responseObj = JSON.parse(localStorage.getItem("responseData"));
    var weight = responseObj.Weight
    var height = responseObj.Height
    var workload = responseObj.WorkLoad
    var gender = responseObj.Gender
    var age = responseObj.Age


    //利用baseUrl加上email才是request Url
    const Url = 'http://127.0.0.1:8080/api/v1/recNutri/calory/?weight=' + weight + "&height=" + height + "?workload=" + workload + "&gender=" + gender + "&age=" + age;


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
                    console.log(responseData)
                    var dataContainer = document.getElementById('dataContainer');
                    //dataContainer.textContent = JSON.stringify(data);
                    var cal = JSON.stringify(responseData.calory);
                    dataContainer.innerHTML = "";

                    //dataContainer.innerHTML = age + select_value+"<br>";
                    dataContainer.innerHTML = "查詢結果 :<br>根據你的身高與體重算出的BMI,你的建議攝取量是<br>"+ cal + "(cal)唷<br>";
                });
            } else {
                // 登入失敗，跳轉到註冊頁面
                console.log("Error")
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });

});

