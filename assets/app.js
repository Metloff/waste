var months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];
var monthsNumber = ['01','02','03','04','05','06','07','08','09','10','11','12'];
var curMonthNumber = new Date().getMonth() + 1;
var curMonthTitle = '';
var labelsBar = [];
var dataBar = [];
var labelsLine = [];
var dataLine = [];

var prepareData = function() {
    // Prepare oneMonth graph.
    prepareMonthData();

    // Prepare oneYear graph.              
    prepareYearData();
};

var prepareMonthData = function() {
  labelsBar = [];
  dataBar = [];
  for(var i = 0; i < window.a[curMonthNumber].Categories.length; i++) {
        labelsBar.push(window.a[curMonthNumber].Categories[i].Category);
        dataBar.push(window.a[curMonthNumber].Categories[i].Amount);
    }  
    curMonthTitle = window.a[curMonthNumber].MonthTitle
};

var prepareYearData = function() {
  for(var i = 1; i <= 12; i++) {
        dataLine[i-1] = window.a[i].Amount;
    }
}

prepareData();
var ctx = document.getElementById("oneMonth");
var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: labelsBar,
        datasets: [{
            label: 'Потрачено денег',
            data: dataBar,
            backgroundColor: [
                'rgba(26, 179, 148, 0.4)',
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(153, 102, 255, 0.2)',
                'rgba(255, 159, 64, 0.2)'
            ],
            borderColor: [
                'rgba(26, 179, 148, 0.7)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(255, 159, 64, 1)'
            ],
            borderWidth: 1
        }]
    },
    options: {
        responsive: true,
        scales: {
            yAxes: [{ ticks: { beginAtZero:true } }]
        },
        legend: {
            position: 'top',
        },
        title: {
            display: true,
            text: 'Статистика за ' + curMonthTitle
        },
        
        onClick: function(e, legendItem) {
            var category_waste = window.a[curMonthNumber].Categories[legendItem[0]._index]
            var ul = document.getElementById('categoryStat');
            var catName = document.getElementById('categoryName');

            // Очищаем список       
            ul.innerHTML = '';
            catName.innerHTML = '';
                      
            // Создаем новый список
            for (var i = 0; i < category_waste.Transactions.length; i++) {
                var li = document.createElement('li');
                li.innerHTML = "<div class='statRow'>" + category_waste.Transactions[i].f2 + ": -" + category_waste.Transactions[i].f1 + "</div>  <div class='gray'>"+ convertDate(category_waste.Transactions[i].f3) + "</div>";
                ul.appendChild(li);
            }
            catName.innerHTML = catName.innerHTML + category_waste.Category.charAt(0).toUpperCase() + category_waste.Category.substring(1).toLowerCase();                 
       },
    }
});

var convertDate = function(unixTime) {
    var date = new Date(unixTime*1000);
    var year = date.getFullYear();
    var month = monthsNumber[date.getMonth()];
    var day = date.getDate();
    return day + "-" + month + "-" + year
};

var ctx2 = document.getElementById("oneYear");
var myLineChart = new Chart(ctx2, {
    type: 'line',
    data: {
        labels: months,
        datasets:[{
            label: "My First Dataset",
            data: dataLine,
            fill: false,
            borderColor: 'rgb(75, 192, 192)',
            lineTension: 0.1}
        ]},
        options: { 
            responsive: true,
            title: {
              display: true,
              text: 'Статистика за послений год'
            },
            onClick: function(e, legendItem) {
              // Prepare oneMonth graph.                  
              curMonthNumber = legendItem[0]._index + 1;
              prepareMonthData();
              myChart.data.datasets[0].data = dataBar;
              myChart.data.labels = labelsBar;
              myChart.options.title.text = 'Статистика за ' + curMonthTitle;
              myChart.update();  
            },
        }
});