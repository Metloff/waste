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
var labels = [];
var data = [];
var prepareData = function() {
    for(var i = 0; i < window.a.length; i++) {
        labels.push(window.a[i].Category);
        data.push(window.a[i].Amount);
    }
}

var ctx = document.getElementById("myChart");
prepareData();
console.log(labels);
var myChart = new Chart(ctx, {
    type: 'bar',
    data: {
        labels: labels,
        datasets: [{
            label: 'Потрачено денег',
            data: data,
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
            console.log(legendItem[0]._index);
            var category_waste = window.a[legendItem[0]._index]
            var ul = document.getElementById('kek');

            // Очищаем список
            for (var i = 0; i < ul.childElementCount; i++) {
                ul.removeChild(ul.children[i]);
            };

            // Создаем новый список
            for (var i = 0; i < category_waste.JSON.length; i++) {
                var li = document.createElement('li');
                li.innerHTML = category_waste.JSON[i].f2 + ": " + category_waste.JSON[i].f1 + "<small class='gray'>Updated at 2017-04-23 19:44 MSK</small>";
                ul.appendChild(li);
            }
        },
    }
});

// Возможно при заполнении списка можно записывать его под ключем в хеш и от туда доставать при повторном нажатии, а не формировать заново