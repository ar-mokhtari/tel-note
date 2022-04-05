function contactHtmlPlacement() {
    //contact
    $("#ContactTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    $("#ContactTabContent div table").append("<thead><tr class='table-primary'>" +
        //header list
        "<td>PersonID</td><td>JobID</td><td>Tel</td><td>Address</td><td>Description</td>" +
        "</tr></thead><tbody></tbody>");
    $.each(AllData.ContactData, function (indexContactData, elementContactData) {
        //body list
        $("#ContactTabContent div table tbody").append("<tr>" +
            "<td>" + (JSON.stringify(elementContactData.PersonID)) + '</td>' +
            "<td>" + (JSON.stringify(elementContactData.JobID)) + '</td>' +
            "<td>" + (JSON.stringify(elementContactData.Tel)) + '</td>' +
            "<td>" + (JSON.stringify(elementContactData.Address)) + '</td>' +
            "<td>" + (JSON.stringify(elementContactData.Description)) + '</td>' +
            "</tr>");
    });
}
