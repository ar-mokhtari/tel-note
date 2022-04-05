const addContactUrl = 'http://127.0.0.1:1212/new-contact';
const addCityUrl = 'http://127.0.0.1:1212/add-city';
const fillDataUrl = 'http://127.0.0.1:1212/fill-data';
const getDataUrl = 'http://127.0.0.1:1212/get-data';
const getMenuUrl = 'http://127.0.0.1:1212/menu-list';
const getReportContactUrl = 'http://127.0.0.1:1212/report-contact';

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

let AllData

function getData() {

    $.ajax({
        type: 'GET',
        url: getDataUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({
            State: "State",
            CityData: "CityData",
            ContactData: "ContactData",
            CustomerData: "CustomerData",
            CustomerGroupData: "CustomerGroupData",
            CustomerGroupRelationData: "CustomerGroupRelationData",
            CustomerRelationData: "CustomerRelationData",
            PersonData: "PersonData",
            CountriesData: "CountriesData",
        }),
        success: function (data) {
            AllData = data
        }
    });
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

function getMenu() {
    $("#MainSidebar").empty();
    $.ajax({
        type: 'GET',
        url: getMenuUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({GroupName: "GroupName", Row: "Row"}),
        Row: JSON.stringify({Title: "Title", Description: "Description", Route: "Route"}),
        success: function (data) {
            $.each(data, function (menuIndex, menuElement) {
                $("#MainSidebar").append(
                    "<li class='mb-1'>" +
                    "<button aria-expanded=\"false\" class=\"btn btn-toggle align-items-center rounded collapsed\"" +
                    " data-bs-target = #group-collapse" + menuIndex + " data-bs-toggle = collapse > " +
                    menuElement.GroupName +
                    "</button>" +
                    "<div class=\"collapse\" id=\"" + "group-collapse" + menuIndex + "\">\n" +
                    "    <ul class=\"btn-toggle-nav list-unstyled fw-normal pb-1 small\">\n" +
                    "    </ul>\n" +
                    "</div>" +
                    "</li>"
                );
                $.each(menuElement.Row, function (detailIndex, detailElement) {
                    $("#group-collapse" + menuIndex + " ul").append("<li><a class=\"link-dark rounded\" href=\"#\">" + detailElement.Description + "</a></li>")
                })
            })
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


//run menu on startup
getMenu()
