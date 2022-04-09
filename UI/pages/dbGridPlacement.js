function dbGridHtmlPlacement() {
    $("#ResponseTabContent div").append("<table data-toggle='table' class='table table-striped table-hover table-responsive small'></table>");
    let tbl_body = document.createElement("tbody");
    let odd_even = false;
    $.each(AllData.Data, function () {
        let tbl_row = tbl_body.insertRow();
        tbl_row.className = odd_even ? "odd" : "even";
        $.each(this, function (k, v) {
            let cell = tbl_row.insertCell();
            if (v != null) {
                cell.appendChild(document.createTextNode(v.toString()));
            }
        });
        odd_even = !odd_even;
    });
    $("#ResponseTabContent div table").append(tbl_body);
}
