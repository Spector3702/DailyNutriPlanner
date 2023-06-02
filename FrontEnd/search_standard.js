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
            <h3>每人每天各種營養素攝取狀況</h3>
            <table>
            <colgroup span="4" style="background-color: #BAA0F2;"></colgroup>
              <tr>
                <td>卡路里</td>
                <td>蛋白質</td>
                <td>脂肪</td>
                <td>糖</td>
              </tr>
              <tr>
                <td>${data.Calorie}(kcal)</td>
                <td>${data.Protein}(g)</td>
                <td>${data.Fat}(g)</td>
                <td>${data.Carbohydrate}(g)</td>
              </tr>
            </table>
            <br>
            <table>
            <colgroup span="7" style="background-color: #BAA0F2;"></colgroup>
              <tr>
                <td>維生素B1</td>
                <td>維生素B2</td>
                <td>維生素C</td>
                <td>維生素B6</td>
                <td>維生素A</td>
                <td>維生素E</td>
                <td>尼古丁</td>
              </tr>
              <tr>
                <td>${data.VitaminB1}(mg)</td>
                <td>${data.VitaminB2}(mg)</td>
                <td>${data.VitaminC}(mg)</td>
                <td>${data.VitaminB6}(mg)</td>
                <td>${data.VitaminA}(mg)</td>
                <td>${data.VitaminE}(mg)</td>
                <td>${data.Nicotine}(mg)</td>
              </tr>
            </table>
            <br>
            <table>
            <colgroup span="7" style="background-color: #BAA0F2;"></colgroup>
              <tr>
                <td>鈣質</td>
                <td>磷</td>
                <td>鐵</td>
                <td>鎂</td>
                <td>鋅</td>
                <td>鈉</td>
                <td>鉀</td>
              </tr>
              <tr>
                <td>${data.Ca}(mg)</td>
                <td>${data.P}(mg)</td>
                <td>${data.Fe}(mg)</td>
                <td>${data.Mg}(mg)</td>
                <td>${data.Zn}(mg)</td>
                <td>${data.Na}(mg)</td>
                <td>${data.K}(mg)</td>
              </tr>
            </table>

          `;
          dataContainer.appendChild(rowElement);
      //test
      //var select = JSON.stringify(data[0]);
      //var Calorie = JSON.stringify(data[0].calorie);
      //var se = JSON.stringify(data[age == "<1"]);
      //以下這段刪掉就不能成功，怪怪的
      /*
      for (var i = 0; i < 18; i++) {
        var age = JSON.stringify(data[i].age);
        var gender = JSON.stringify(data[i].gender);
      }*/
      //dataContainer.innerHTML = age + select_value+"<br>";
      //dataContainer.innerHTML = "查詢結果" + select+"<br>"+Calorie+select_value;
      
    })

    /*catch*/
    .catch(error => {
      console.log("錯誤：" + error);
    });
});
  /*
"Gender":          gender,
"Age":             age,
"Calorie":         recommendedNutrition.Calorie,
"Protein":         recommendedNutrition.Protein,
"Fat":             recommendedNutrition.Fat,
"Carbohydrate":    recommendedNutrition.Carbohydrate,
"VitaminB1":       recommendedNutrition.VitaminB1,
"VitaminB2":       recommendedNutrition.VitaminB2,
"VitaminC":        recommendedNutrition.VitaminC,
"Nicotine":        recommendedNutrition.Nicotine,
"VitaminB6":       recommendedNutrition.VitaminB6,
"VitaminA":        recommendedNutrition.VitaminA,
"VitaminE":        recommendedNutrition.VitaminE,
"Ca":              recommendedNutrition.Ca,
"P":               recommendedNutrition.P,
"Fe":              recommendedNutrition.Fe,
"Mg":              recommendedNutrition.Mg,
"Zn":              recommendedNutrition.Zn,
"Na":              recommendedNutrition.Na,
"K":               recommendedNutrition.K,
"NumbersOfPeople":
 
<p>卡路里:${row.calorie}<p>
<p>蛋白質：${row.protein}</p>
<p>脂肪：${row.fat}</p>
<p>糖：${row.carbohydrate}</p>
<p>維生素B1:${row.vitaminB1}</p>
<p>維生素B2:${row.vitaminB2}</p>
<p>維生素C:${row.vitaminC}</p>
<p>維生素C:${row.vitaminB6}</p>
<p>維生素C:${row.vitaminA}</p>
<p>維生素C:${row.vitaminE}</p>
<p>維生素C:${row.Nicotine}</p>
<p>鈣質:${row.ca}</p>
<p>鈣質:${row.p}</p>
<p>鈣質:${row.fe}</p>
<p>鈣質:${row.mg}</p>
<p>鈣質:${row.zn}</p>
<p>鈣質:${row.na}</p>
<p>鈣質:${row.k}</p>
*/