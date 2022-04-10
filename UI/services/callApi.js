const hostName = 'http://127.0.0.1:1212'
const addContactUrl = hostName + '/new-contact';
const addCityUrl = hostName + '/add-city';
const fillDataUrl = hostName + '/fill-data';
const getDataUrl = hostName + '/get-data';
const getMenuUrl = hostName + '/menu-list';
const getReportContactUrl = hostName + '/report-contact';

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


function getDataAndApplyCatch() {
    $("#ResponseTabContent div," +
        "#ContactTabContent div, " +
        "#CustomerTabContent div").empty();
    //city
    cityHtmlPlacement();
    //contact
    contactHtmlPlacement();
    // //CustomerData
    customerHtmlPlacement();
    //header and state
    $("#statusRespond").empty().append("<span>" + JSON.stringify(AllData.State) + '</span>');
}

function getReportContact() {
    $("#ReportContactTabContent div").empty();
    $.ajax({
        type: 'GET',
        url: getReportContactUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({
            State: "State",
            ReportContactData: "ReportContactData",
        }),
        success: function (data) {
            //ReportContact
            contactReportPlacement(data);
            //header and state
            $("#statusRespond").empty().append("<span>" + JSON.stringify(data.State) + '</span>');
        }
    });
}

function fillData() {
    $.ajax({
        type: 'POST',
        url: fillDataUrl,
        contentType: "application/json; charset=utf-8",
        success: function (data) {
            getData()
            alert("Call API for load data successful")
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
            getDataAndApplyCatch()
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
            getDataAndApplyCatch()
        }
    });
}

$('.getDataBtn').click(function () {
    getData()
})
$('.applyBtn').click(function () {
    getDataAndApplyCatch()
})
$('.getReportContactBtn').click(function () {
    getReportContact()
})
$('.fillDataBtn').click(function () {
    fillData()
})
$('.addCityBtn').click(function () {
    addCity()
})
$('.addContactBtn').click(function () {
    addContact()
})
