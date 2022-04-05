function contactReportPlacement(data) {
    $("#ReportContactTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    $("#ReportContactTabContent div table").append("<thead><tr class='table-primary'>" +
        "<td>ID</td><td>PID</td> <td>PersonFullName</td><td>DOB</td><td>JID</td><td>JobName</td><td>Gender</td><td>Cellphone</td><td>LoID</td><td>JobCityName</td><td>Address</td><td>Description</td>" +
        "</tr></thead><tbody></tbody>");
    $.each(data.ReportContactData, function (index, element) {
        $("#ReportContactTabContent div table tbody").append("<tr>" +
            "<td>" + (JSON.stringify(element.ID)) + '</td>' +
            "<td>" + (JSON.stringify(element.PID)) + '</td>' +
            "<td>" + ((JSON.stringify(element.PersonName)) + " " + (JSON.stringify(element.PersonFamily))).replace(/"/g, '') + '</td>' +
            "<td>" + mask(((JSON.stringify(element.DOB).replace(/"/g, '')).replace(/-/g, '')), "**** ** **") + '</td>' +
            "<td>" + (JSON.stringify(element.JID)) + '</td>' +
            "<td>" + (JSON.stringify(element.JobName)) + '</td>' +
            "<td>" + (JSON.stringify(element.Gender)) + '</td>' +
            "<td>" + (JSON.stringify(element.Cellphone)) + '</td>' +
            "<td>" + (JSON.stringify(element.LoID)) + '</td>' +
            "<td>" + (JSON.stringify(element.JobCityName)) + '</td>' +
            "<td>" + (JSON.stringify(element.Address)) + '</td>' +
            "<td>" + (JSON.stringify(element.Description)) + '</td>' +
            "</tr>");
    });
}
