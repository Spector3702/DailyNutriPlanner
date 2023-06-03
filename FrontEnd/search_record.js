document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的email文字
    var date = document.getElementById('date').value;
    const responseObj = JSON.parse(localStorage.getItem("responseData"));
    console.log(date);
    //利用baseUrl加上email才是request Url
    const Url = 'http://127.0.0.1:8080/api/v1/record/daily/?date=' + date + "&email="+responseObj.email;


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
            var food = responseData.records;
            //var test = JSON.stringify(food[0].Name);
            dataContainer.innerHTML = "";
            dataContainer.innerHTML = "查詢成功"+"<br>"+"你在"+date+"紀錄的食物有:";
            for(i=0;i<responseData.toString().length;i++){
                var rowElement = document.createElement("div");
                rowElement.innerHTML = `
                <p>食物名稱 : ${food[i].Name}<p>
                `;
                dataContainer.appendChild(rowElement);
            }
        })
        .catch(error => {
            console.error('Error:', error);
        });

});

