const addContactUrl = 'http://127.0.0.1:1212/new-contact';
const addCityUrl = 'http://127.0.0.1:1212/add-city';
const fillDataUrl = 'http://127.0.0.1:1212/fill-data';
const getDataUrl = 'http://127.0.0.1:1212/get-data';
const getMenuUrl = 'http://127.0.0.1:1212/menu-list';

const newCityData = {
    "Name": "شهر جدید",
    "EnglishName": "new town",
    "AriaCode": "131313",
    "Lat": "36.5678161",
    "Lng": "51.8789308"
}
const newContactData = {
    "PersonID": 1,
    "JobID": 1,
    "Tel": "09121212",
    "Address": "none",
    "Description": "---",
}

function getData() {
    $("#response span").empty();
    $.ajax({
        type: 'GET',
        url: getDataUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({CityData: "CityData", ContactData: "ContactData"}),
        success: function (data) {
            $("#response span").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
            $("#response span table").append("<thead><tr class='table-primary'><td>Name</td><td>EnglishName</td><td>AriaCode</td><td>Lat</td><td>Lng</td></tr></thead><tbody></tbody>");
            $.each(data.CityData, function (index, element) {
                $("#response span table tbody").append("<tr>" + "<td>" + (JSON.stringify(element.Name)) + '</td>' +
                    "<td>" + (JSON.stringify(element.EnglishName)) + '</td>' +
                    "<td>" + (JSON.stringify(element.AriaCode)) + '</td>' +
                    "<td>" + (JSON.stringify(element.Lat)) + '</td>' +
                    "<td>" + (JSON.stringify(element.Lng)) + '</td>' +
                    "</tr>");
            });
            // $("#response span").append("<span>" + JSON.stringify(data.ContactData) + '</span>');
        }
    });
}

function fillData() {
    $.ajax({
        type: 'POST',
        url: fillDataUrl,
        contentType: "application/json; charset=utf-8",
        success: function (data) {
            alert("OK")
        }
    });
}

function addCity() {
    $.ajax({
        type: 'POST',
        url: addCityUrl,
        data: JSON.stringify(newCityData),
        contentType: "application/json; charset=utf-8",
        traditional: true,
        success: function (data) {
            getData()
        }
    });
}

function addContact() {
    $.ajax({
        type: 'POST',
        url: addContactUrl,
        data: JSON.stringify(newContactData),
        contentType: "application/json; charset=utf-8",
        // traditional: true,
        success: function (data) {
        }
    });
}

function getMenu() {
    $.ajax({
        type: 'GET',
        url: getMenuUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({GroupName: "GroupName", Row: "Row"}),
        Row: JSON.stringify({Title: "Title", Description: "Description"}),
        success: function (data) {
            $("#MainSidebar ul").empty()
            $.each(data, function (index, element) {
                $("#MainSidebar ul").append(
                    "<li class='mb-1'>" +
                    "<button class=\"btn btn-toggle align-items-center rounded collapsed\" data-bs-toggle=\"collapse\"\n" +
                    "data-bs-target=\"" + "#group-collapse" + index + "\" aria-expanded=\"false\">\n" +
                    element.GroupName +
                    "</button>" +
                    "<div class=\"collapse\" id=\"" + "group-collapse" + index + "\">\n" +
                    "    <ul class=\"btn-toggle-nav list-unstyled fw-normal pb-1 small\">\n" +
                    "<li><a href=\"#\" class=\"link-dark rounded\">" + JSON.stringify(element.Row) + "</a></li>" +
                    // $.each(element.Row, function (detailIndex, detailElement) {
                    //     console.log(detailElement.Description)
                    //     // alert(JSON.stringify(detailElement.Description))
                    //     // this.append("<li><a href=\"#\" class=\"link-dark rounded\">" + JSON.stringify(detailElement.Description) + "</a></li>")
                    // })
                    "    </ul>\n" +
                    "</div>" +
                    "</li>" +
                    "</ul>"
                )
            });
        }
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
$('.addContactBtn').click(function () {
    // addContact()
    getMenu()
})

