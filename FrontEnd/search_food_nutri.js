document.querySelector('form').addEventListener('submit', function (event) {
    event.preventDefault();

    //利用getElementById抓取在form中的文字
    var food_bar = document.getElementById('select_food').value;

    //建一個變數來儲存剛剛的那些資料
    var formData = new URLSearchParams();
    formData.append('name', food_bar);

    const Url = 'http://127.0.0.1:8080/api/v1/everydayNutri/GetNutri/?name=' + food_bar;

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

            //顯示資料到網頁
            var dataContainer = document.getElementById('dataContainer');

            //dataContainer.textContent = JSON.stringify(data);
            var select = JSON.stringify(responseData[0].analysis_item);
            dataContainer.innerHTML = "";
            var rowElement = document.createElement("div");
            rowElement.innerHTML = `
                    <h3>成功查詢</h3>
                    <table>
                        <thead>
                            <th>${responseData[0].analysis_item}</th>
                            <th>${responseData[1].analysis_item}</th>
                            <th>${responseData[2].analysis_item}</th>
                            <th>${responseData[3].analysis_item}</th>
                            <th>${responseData[4].analysis_item}</th>
                            <th>${responseData[5].analysis_item}</th>
                            <th>${responseData[6].analysis_item}</th>
                        </thead>
                        <tbody>
                            <tr>
                                <td>${responseData[0].content_per_unit}(${responseData[0].unit})</td>
                                <td>${responseData[1].content_per_unit}(${responseData[1].unit})</td>
                                <td>${responseData[2].content_per_unit}(${responseData[2].unit})</td>
                                <td>${responseData[3].content_per_unit}(${responseData[3].unit})</td>
                                <td>${responseData[4].content_per_unit}(${responseData[4].unit})</td>
                                <td>${responseData[5].content_per_unit}(${responseData[5].unit})</td>
                                <td>${responseData[6].content_per_unit}(${responseData[6].unit})</td>
                            </tr>
                        </tbody>
                    </table>
                    <br>
                    <table>       
                        <thead>
                            <th>${responseData[7].analysis_item}</th>
                            <th>${responseData[8].analysis_item}</th>
                            <th>${responseData[9].analysis_item}</th>
                            <th>${responseData[10].analysis_item}</th>
                            <th>${responseData[11].analysis_item}</th>
                            <th>${responseData[12].analysis_item}</th>
                        </thead>
                        <tbody>
                        <tr>
                            <td>${responseData[7].content_per_unit}(${responseData[7].unit})</td>
                            <td>${responseData[8].content_per_unit}(${responseData[8].unit})</td>
                            <td>${responseData[9].content_per_unit}(${responseData[9].unit})</td>
                            <td>${responseData[10].content_per_unit}(${responseData[10].unit})</td>
                            <td>${responseData[11].content_per_unit}(${responseData[11].unit})</td>
                            <td>${responseData[12].content_per_unit}(${responseData[12].unit})</td>
                        </tr>
                        </tbody>
                    </table>
                    <br>
                    <table>
                        <thead>
                            <th>${responseData[13].analysis_item}</th>
                            <th>${responseData[14].analysis_item}</th>
                            <th>${responseData[15].analysis_item}</th>
                            <th>${responseData[16].analysis_item}</th>
                            <th>${responseData[17].analysis_item}</th>
                        </thead>
                        <tbody>
                        <tr>
                            <td>${responseData[13].content_per_unit}(${responseData[13].unit})</td>
                            <td>${responseData[14].content_per_unit}(${responseData[14].unit})</td>
                            <td>${responseData[15].content_per_unit}(${responseData[15].unit})</td>
                            <td>${responseData[16].content_per_unit}(${responseData[16].unit})</td>
                            <td>${responseData[17].content_per_unit}(${responseData[17].unit})</td>
                        </tr>
                        </tbody>
                    </table>
                    `;
            dataContainer.appendChild(rowElement);
            
        })//end of .then
        .catch(error => {
            console.error('Error:', error);
            var dataContainer = document.getElementById('dataContainer');

            dataContainer.innerHTML = "";
            var rowElement = document.createElement("div");
              rowElement.innerHTML = `       
              <h3>查詢失敗<br>請檢查輸入的名稱是否錯誤</h3>  
              `;
            dataContainer.appendChild(rowElement);
        });

});
