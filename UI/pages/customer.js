function customerHtmlPlacement() {
    //customerHtmlPlacement
    $("#CustomerTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    $("#CustomerTabContent div table").append("<thead><tr class='table-primary'>" +
        //header list
        "<td>PersonID</td><td>Description</td><td>CreateAt</td><td>UpdatedAt</td>" +
        "</tr></thead><tbody></tbody>");
    $.each(AllData.CustomerData, function (indexCustomerData, elementCustomerData) {
        //body list
        $("#CustomerTabContent div table tbody").append("<tr>" +
            "<td>" + (JSON.stringify(elementCustomerData.PersonID)) + '</td>' +
            "<td>" + (JSON.stringify(elementCustomerData.Description)) + '</td>' +
            "<td>" + mask(((JSON.stringify(elementCustomerData.CreateAt).replace(/"/g, '')).replace(/-/g, '')), "**** ** **") + '</td>' +
            "<td>" + mask(((JSON.stringify(elementCustomerData.UpdatedAt).replace(/"/g, '')).replace(/-/g, '')), "**** ** **") + '</td>' +
            "</tr>");
    });
}
