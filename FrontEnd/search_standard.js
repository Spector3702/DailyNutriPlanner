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
function fetchData() {
  // 獲取選取的欄位值
  var selectElement = document.getElementById("select_age");
  var select_value = selectElement.value;
  // 發出API請求
  fetch('http://localhost:8080/api/v1/recNutri/?age='+select_value)
    .then(response => response.json())
    .then(data => {
      console.log(data);
      // 将获取到的数据显示在 HTML 中
      var dataContainer = document.getElementById('dataContainer');
      //dataContainer.textContent = JSON.stringify(data[0]);
      
      dataContainer.innerHTML = "";
      //test
      var select = JSON.stringify(data[0]);
      var Calorie = JSON.stringify(data[0].calorie);
      var se = JSON.stringify(data[age=="<1"]);
      //以下這段刪掉就不能成功，怪怪的
      for(var i=0;i<18;i++){
        var age = JSON.stringify(data[i].age);
      }
      //dataContainer.innerHTML = age + select_value+"<br>";
      //dataContainer.innerHTML = "查詢結果" + select+"<br>"+Calorie+select_value;
      data.forEach(row => {
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
      });
    })


    .catch(error => {
      console.log("錯誤：" + error);
    });
  }
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
		"NumbersOfPeople":*/