var lineOptions = {
	// animation: false,
	// scaleShowGridLines: true,
	// scaleGridLineColor: "rgba(0,0,0,.05)",
	// scaleGridLineWidth: 1,
	// pointDot: true,
	// pointDotRadius: 3.5,
	// pointDotStrokeWidth: 1,
	// pointHitDetectionRadius: 10,
	// datasetStroke: true,
	// datasetStrokeWidth: 2,
	// datasetFill: true,
	// responsive: true,
	// maintainAspectRatio: false,
	// legend: {
	// 	display: false
	// },
	// tooltips: {
	// 	mode: "label"
	// }
};
var months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'June', 'July', 'Aug', 'Sept', 'Oct', 'Nov', 'Dec'];
var monthsNumber = ['01','02','03','04','05','06','07','08','09','10','11','12'];
var labelsBar = [];
var dataBar = [];
var labelsLine = [];
var dataLine = [];

var prepareData = function() {
    for(var i = 0; i < window.a.CurMonth.length; i++) {
        labelsBar.push(window.a.CurMonth[i].Category);
        dataBar.push(window.a.CurMonth[i].Amount);
    }  

    for(var i = 1; i <= 12; i++) {        
        for(var j = 0; j < window.a.YearStat.length; j++) {
            console.log(window.a.YearStat[j].Month);
            if(window.a.YearStat[j].Month == i) {
               dataLine.push(window.a.YearStat[j].Amount);
               return 
            }
            dataLine.push(0);  
        }
    }
};

var ctx = document.getElementById("myChart");
prepareData();
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
            text: 'Chart.js Bar Chart'
        },
        
		onClick: function(e, legendItem) {
            var category_waste = window.a.CurMonth[legendItem[0]._index]
            var ul = document.getElementById('categoryStat');
            var catName = document.getElementById('categoryName');

            // Очищаем список       
            ul.innerHTML = '';
            catName.innerHTML = '';
                      
            // Создаем новый список
            for (var i = 0; i < category_waste.JSON.length; i++) {
                var li = document.createElement('li');
                li.innerHTML = "<div class='statRow'>" + category_waste.JSON[i].f2 + ": -" + category_waste.JSON[i].f1 + "</div>  <div class='gray'>"+ convertDate(category_waste.JSON[i].f3) + "</div>";
                ul.appendChild(li);
            }
            catName.innerHTML = catName.innerHTML + category_waste.Category;  
            console.log(category_waste);                   
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

console.log(labelsLine);
console.log(window.a.YearStat[0].Month);

var ctx2 = document.getElementById("myChart2");
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
        options: { responsive: true}
});