//查詢食物營養成分的js
/*var age = document.getElementById("age").value;

fetch('http://localhost:8080/api/v1/recNutri/')
  .then(response => response.json())
  .then(data => {
    // 在这里处理从URL获取的数据
    // 可以更新HTML元素来显示数据
    console.log(data);
    const dataContainer = document.getElementById('dataContainer');
    dataContainer.textContent = JSON.stringify(data);
 
  })

  .catch(error => {
    console.error('Error:', error);
  });
*/
document.querySelector('form').addEventListener('submit', function (event) {
  event.preventDefault();
  // 獲取選取的欄位值
  var age = document.getElementById("select_age").value;
  var gender = document.getElementById("select_gender").value;

  const Url = 'http://localhost:8080/api/v1/recNutri/nutri/?gender=' + gender + '&age=' + age
  fetch(Url, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
    }
  })
    .then(response => response.json())
    .then(data => {
      console.log(data);
      // 將資料顯示在網頁上
      var dataContainer = document.getElementById('dataContainer');
      //dataContainer.textContent = JSON.stringify(data);

      dataContainer.innerHTML = "";
      var rowElement = document.createElement("div");
          rowElement.innerHTML = `
            <h3>成功查詢</h3>
            <h3>每人每天各種營養素攝取狀況:</h3>
            <table>
            <thead>
                <th>卡路里</th>
                <th>蛋白質</th>
                <th>脂肪</th>
                <th>糖</th>
            </thead>
            <tbody>
              <tr>
                <td>${data.Calorie}(kcal)</td>
                <td>${data.Protein}(g)</td>
                <td>${data.Fat}(g)</td>
                <td>${data.Carbohydrate}(g)</td>
              </tr>
            </tbody>
            </table>
            <br>
            <table>
            <thead>
                <th>維生素B1</th>
                <th>維生素B2</th>
                <th>維生素C</th>
                <th>維生素B6</th>
                <th>維生素A</th>
                <th>維生素E</th>
                <th>尼古丁</th>
            </thead>
            <tbody>
              <tr>
                <td>${data.VitaminB1}(mg)</td>
                <td>${data.VitaminB2}(mg)</td>
                <td>${data.VitaminC}(mg)</td>
                <td>${data.VitaminB6}(mg)</td>
                <td>${data.VitaminA}(mg)</td>
                <td>${data.VitaminE}(mg)</td>
                <td>${data.Nicotine}(mg)</td>
              </tr>
            </tbody>
            </table>
            <br>
            <table>
            <thead>
                <th>鈣質</th>
                <th>磷</th>
                <th>鐵</th>
                <th>鎂</th>
                <th>鋅</th>
                <th>鈉</th>
                <th>鉀</th>
            </thead>
            <tbody>
              <tr>
                <td>${data.Ca}(mg)</td>
                <td>${data.P}(mg)</td>
                <td>${data.Fe}(mg)</td>
                <td>${data.Mg}(mg)</td>
                <td>${data.Zn}(mg)</td>
                <td>${data.Na}(mg)</td>
                <td>${data.K}(mg)</td>
              </tr>
            </tbody>
            </table>
          `;
          dataContainer.appendChild(rowElement);
    })

    /*catch*/
    .catch(error => {
      console.log("錯誤：" + error);
    });
});
