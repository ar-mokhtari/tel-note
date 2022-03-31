const addCityUrl = 'http://127.0.0.1:1212/add-city';
const fillDataUrl = 'http://127.0.0.1:1212/fill-data';
const getDataUrl = 'http://127.0.0.1:1212/get-data';
const newCityData = {
    "Name": "شهر قل قل",
    "EnglishName": "ghol ghol town",
    "AriaCode": "131313",
    "Lat": "36.5678161",
    "Lng": "51.8789308"
}

function getData() {
    $.ajax({
        type: 'GET',
        url: getDataUrl,
        contentType: "application/x-www-form-urlencoded; charset=iso-8859-1",
        dataType: "json",
        success: function (data) {
            let cityData = JSON.stringify(data.CityData);
            $.each(JSON.stringify(cityData), function (index, element) {
                $("#response span").append("<span>" + (JSON.stringify(element.Name)) + '</span>');
                $("#response span").append("<span>" + (JSON.stringify(element.EnglishName)) + '</span>');
                $("#response span").append("<span>" + (JSON.stringify(element.AriaCode)) + '</span>');
                $("#response span").append("<span>" + (JSON.stringify(element.Lat)) + '</span>');
                $("#response span").append("<span>" + (JSON.stringify(element.Lng)) + '</span>');
            });
        }
    });
}

function fillData() {
    $.post(fillDataUrl, {}, function (data, status) {
        console.log('${data} and status is ${status}')
    });
}

$('.getDataBtn').click(function () {
    getData()
})
$('.fillDataBtn').click(function () {
    fillData()
})
$('.addCityBtn').click(function () {
    addCity()
})

function addCity() {
    $.ajax({
        type: 'POST',
        url: addCityUrl,
        data: JSON.stringify(newCityData),
        // contentType: "application/json; charset=utf-8",
        traditional: true,
        success: function (data) {
            alert("123")
        }
    });
}

