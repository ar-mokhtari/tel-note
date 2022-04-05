function cityHtmlPlacement() {
    $("#ResponseTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    $("#ResponseTabContent div table").append("<thead><tr class='table-primary'>" +
        //header list
        "<td>Name</td><td>EnglishName</td><td>AriaCode</td><td>Lat</td><td>Lng</td>" +
        "</tr></thead><tbody></tbody>");
    $.each(AllData.CityData, function (index, element) {
        //body list
        $("#ResponseTabContent div table tbody").append("<tr>" +
            "<td>" + (JSON.stringify(element.Name)).replace(/"/g, '') + '</td>' +
            "<td>" + (JSON.stringify(element.EnglishName)) + '</td>' +
            "<td>" + (JSON.stringify(element.AriaCode)) + '</td>' +
            "<td>" + mask((JSON.stringify(element.Lat)), "** * **") + '</td>' +
            "<td>" + mask((JSON.stringify(element.Lng)), "** * **") + '</td>' +
            "</tr>");
    });
}
