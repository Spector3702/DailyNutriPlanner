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
    var name = document.getElementById("select_food").value;
    // 發出API請求
    const Url = 'http://localhost:8080/api/v1/foodstuff/name/?name=' + name;

    fetch(Url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded; charset=utf-8'
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            // 将获取到的数据显示在 HTML 中
            var dataContainer = document.getElementById('dataContainer');
            //dataContainer.textContent = JSON.stringify(data);
            var select = JSON.stringify(data);
            dataContainer.innerHTML = "";
            //test
            //dataContainer.innerHTML = "查詢結果" + select+"<br>";
            //成功查詢的表格
            var rowElement = document.createElement("div");
            rowElement.innerHTML = `
            
              <h3>成功查詢</h3>
              <table>
              <colgroup span="4" style="background-color: #BAA0F2;"></colgroup>
              <tr>
                <td>食物成稱</td>
                <td>別名</td>
                <td>英文名稱</td>
                <td>食物狀況</td>
              </tr>
              <tr>
                <td>${data.name}</td>
                <td>${data.other_name}</td>
                <td>${data.english_name}</td>
                <td>${data.describe}</td>
              </tr>
              </table>
              `;
            dataContainer.appendChild(rowElement);
            /*data.forEach(row => {
              if (row.age === select_value) {
                var rowElement = document.createElement("div");
                rowElement.innerHTML = `
                  <h3>成功查詢</h3>
                  <p>卡路里:${row.calorie}<p>
                  <p>脂肪：${row.fat}</p>
                  <p>糖：${row.carbohydrate}</p>
                  <p>維生素B1:${row.vitaminB1}</p>
                  <p>維生素C:${row.vitaminC}</p>
                  <p>鈣質:${row.ca}</p>
         
                `;
                dataContainer.appendChild(rowElement);
              }
            });*/
        })  //end of second .then

        //catch
        .catch(error => {
            console.log("錯誤：" + error);
        });
});
