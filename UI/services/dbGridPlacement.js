//fill data in grid
function dbGridHtmlPlacement() {
    $("#ResponseTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    let tbl_head = document.createElement("thead");
    let tbl_body = document.createElement("tbody");
    let odd_even = false;
    let tbl_rowHead = tbl_head.insertRow();
    $.each(AllData.Data[0], function (index, element) {
        let cell = tbl_rowHead.insertCell();
        cell.appendChild(document.createTextNode(index.toString()));
    });
    $("#ResponseTabContent div table").append(tbl_head);

    $.each(AllData.Data, function () {
        let tbl_row = tbl_body.insertRow();
        tbl_row.className = odd_even ? "odd" : "even";
        $.each(this, function (k, v) {
            let cell = tbl_row.insertCell();
            cell.appendChild(document.createTextNode("" + v));
        });
        odd_even = !odd_even;
    });
    $("#ResponseTabContent div table").append(tbl_body);
    $("#statusRespond").empty().append("<span> " + AllData.State + " </span>");
}
